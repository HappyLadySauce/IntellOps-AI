package gform

// LockUpdate sets the lock for update for current operation.
func (m *Model) LockUpdate() *Model {
	model := m.getModel()
	model.lockInfo = "FOR UPDATE"
	return model
}

// LockShared sets the lock in share mode for current operation.
func (m *Model) LockShared() *Model {
	model := m.getModel()
	model.lockInfo = "LOCK IN SHARE MODE"
	return model
}
