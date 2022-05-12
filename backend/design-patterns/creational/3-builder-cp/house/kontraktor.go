package house

// nah ketika semua sudah kita implementasikan kita bisa membuat Kontraktor dimana kontraktor tersebut yang akan menentukan builder nya
type kontraktor struct {
	builder HouseBuilder
}

func NewKontraktor(builder HouseBuilder) *kontraktor {
	return &kontraktor{
		// TODO: answer here
		builder: builder,
	}
}

func (d *kontraktor) BuildHouse() House {
	// TODO: answer here
	d.builder.buildWindow(5)
	d.builder.buildDoor()
	d.builder.buildGarage()
	d.builder.buildSwimmingPool()

	return d.builder.getHouse()
}

// Rumah yang tidak memiliki SwimmingPool hanya boleh mimiliki satu jendela
func (d *kontraktor) BuildHouseWithoutSwimmingPool() House {
	// TODO: answer here
	d.builder.buildWindow(1)
	d.builder.buildDoor()
	d.builder.buildGarage()

	return d.builder.getHouse()
}
