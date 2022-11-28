package usb

/*
	 Функция получает список блочных устройств и определяет их статус:
	    - mount_RO
		- mount_RW
		- unmount
		- skip
*/
func GetUSBDevices() (res map[string]string, err error) {

	return nil, nil
}

/* Функция для монтирования флешки только для чтения
 */
func MountUSBRO(device string, target string) (err error) {

	return nil
}

/* Функция для монтирования флешки для чтения и записи
 */
func MountUSBRW(device string, target string) (err error) {

	return nil
}

/* Функция для размонтирования устройства
 */
func UmountUSB(target string) (err error) {

	return nil
}

/* Функция для перемонтирования устройства в режим для записи и чтения
 */
func RemountUSBtoRW(device string) (err error) {

	return nil
}

/* Функция для монтирования устройства
 */
func mountDev(device string, target string, mode string) (err error) {

	return nil
}

/* Функция для перемонтирования устройства
 */
func remountDev(device, target, newMode string) (err error) {

	return nil
}

/* Функция для размонтирования устройств
 */
func umountDev(device string) (err error) {

	return nil
}
