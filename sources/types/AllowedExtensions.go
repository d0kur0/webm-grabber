package types

type AllowedExtensions []string

func (allowedExtensions AllowedExtensions) Contains(desiredExtension string) (result bool) {
	result = false

	for _, extension := range allowedExtensions {
		if extension == desiredExtension {
			result = true
			break
		}
	}

	return
}
