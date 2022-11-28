package state

type ACULaState struct {
	Current map[string]string
}

/* Метод для обновления статуса устройства
 */
func (s *ACULaState) UpdateState(device string, status string) (err error) {

	return nil
}
