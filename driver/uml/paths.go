package uml

import "path/filepath"

func (m *Machine) mDir() string {
	return filepath.Join(m.ud.RunDir, "machine", m.Id())
}

func (m *Machine) nsDir() string {
	return filepath.Join(m.ud.RunDir, "ns", m.namespace, m.name)
}

func (m *Machine) umlDir() string {
	return filepath.Join(m.ud.RunDir, "machine", m.Id(), m.Id())
}

func (m *Machine) diskPath() string {
	return filepath.Join(m.ud.StorageDir, "overlay", m.Id()+".disk")
}
