package storage

  type Point struct {
	X float32
	Y float32 
  }
  type Bike struct{
	ID int		`json:"id"`
	Loc Point 
  }


  var AllBikes = map[int]Bike{
	  	1:Bike{
			ID:1,
			Loc: Point{
				X:48.677,
				Y:2.245,
			},
		},
		2:Bike{
			ID:2,
			Loc: Point{
				X:43.3344,
				Y:2.2344,
			},
		},
  }


  func ReturnBikeByID(id int ) Bike{
	  if key,ok := AllBikes[id];ok{
		  return key
	  }
	return Bike{}
  }

  func ReturnAllBikes(response chan []Bike ){
	result :=  []Bike{
		Bike{
			ID:1,
			Loc: Point{
				X:48.677,
				Y:2.245,
			},
		},
		Bike{
			ID:2,
			Loc: Point{
				X:43.3344,
				Y:2.2344,
			},
		},
	}

	response <- result
  } 