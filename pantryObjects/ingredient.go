package pantryObjects

type Ingredient struct {
    Name          string
    Perishable    bool
    ExpDate       string // should this be some date-time object?
    Tags          []Tag
}

type Tag struct {
    Title         string
}

func (i Ingredient) GetName() string {
    return i.Name
}
