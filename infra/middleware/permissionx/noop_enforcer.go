package permissionx

import "log"

var _ Enforcer = &NoopEnforcer{}

// NoopEnforcer 允许所有操作（无权限服务时的默认实现）
type NoopEnforcer struct{}

func NewNoopEnforcer() *NoopEnforcer {
	return &NoopEnforcer{}
}

func (m *NoopEnforcer) ReloadPolicy() error {
	return m.LoadPolicy()
}

func (m *NoopEnforcer) LoadPolicy() error {
	log.Println("[Perm] NoopEnforcer LoadPolicy")
	return nil
}

func (m *NoopEnforcer) Enforce(user string, resource string, action string) (bool, error) {
	return true, nil
}
