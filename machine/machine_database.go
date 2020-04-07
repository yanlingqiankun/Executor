package machine


import (
	"encoding/json"
	"github.com/yanlingqiankun/Executor/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)


type machineDB struct {
	sync.Map
}

type dbItem struct {
	machine   *Base
	StoragePath string `json:"storagePath"`
}

func (db *machineDB) init() {
	logger.Debug("load machine database from disk")
	if !util.PathExist(machineRootDir) {
		err := os.MkdirAll(machineRootDir, 0600)
		if err != nil {
			logger.WithError(err).Error("failed to create machine root dir")
			return
		}
	}
	data, err := ioutil.ReadFile(filepath.Join(machineRootDir, "machines.json"))
	if err != nil {
		if os.IsExist(err) {
			logger.WithError(err).Error("failed to load machine database")
		} else {
			logger.WithError(err).Warning("machine database doesn't exist")
		}
	} else {
		tempDB := make(map[string]*dbItem)
		counts := 0
		if err = json.Unmarshal(data, &tempDB); err != nil {
			logger.WithError(err).Error("failed to parse machine database")
		}
		for _, item := range tempDB {
			path := item.StoragePath
			machine := recovermachineFromFile(path)

			if machine != nil {
				db.add(machine, path)
				counts += 1
			}
		}
		logger.Printf("found %d machines", counts)
	}
}


func recovermachineFromFile(path string) *Base {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	machine := &Base{}
	if err = json.Unmarshal(data, machine); err != nil {
		return nil
	}
	return machine
}


func (db *machineDB) add(machine *Base, path string) {
	db.Store(machine.ID, &dbItem{
		machine:   machine,
		StoragePath: path,
	})
}

func (db *machineDB) save(saveItem bool, id string) {
	logger.Debug("saving machines database")
	data, _ := json.Marshal(db.convertToMap())
	err := ioutil.WriteFile(filepath.Join(machineRootDir, "machines.json"), data, 0600)
	if err != nil {
		logger.WithError(err).Error("cannot save machines database file to disk")
	}

	if saveItem {
		db.saveItem(id)
	}
}

func (db *machineDB) getItem(id string) (*dbItem, bool) {
	if value, ok := db.Load(id); ok {
		item := value.(*dbItem)
		return item, true
	}
	return nil, false
}

func (db *machineDB) convertToMap() map[string]*dbItem {
	m := make(map[string]*dbItem)
	db.Range(func(k, v interface{}) bool {
		id, item := k.(string), v.(*dbItem)
		m[id] = item
		return true
	})
	return m
}

func (db *machineDB) saveItem(id string) {
	logger.Debug("saving a machine")
	if !util.PathExist(filepath.Join(machineRootDir, id)) {
		_ = os.MkdirAll(filepath.Join(machineRootDir, id), 0600)
	}
	if item, ok := db.getItem(id); ok {
		data, _ := json.Marshal(item.machine)
		err := ioutil.WriteFile(item.StoragePath, data, 0600)
		if err != nil {
			logger.WithError(err).Error("cannot save machine database file to disk")
		}
	}

}

func (db *machineDB) delete(id string) {
	db.Delete(id)
}
