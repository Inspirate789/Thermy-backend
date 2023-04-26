package interfaces

// DTO - Data Transfer Object

type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ModelsIdDTO struct {
	Models []int `json:"models_id,omitempty"`
}

type ModelNamesDTO struct {
	Models []string `json:"model_names,omitempty"`
}

type OutputModelDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type ModelElementsIdDTO struct {
	ModelElements []int `json:"model_elements_id,omitempty"`
}

type ModelElementNamesDTO struct {
	ModelElements []string `json:"model_element_names,omitempty"`
}

type OutputModelsDTO struct {
	Models []OutputModelDTO `json:"models,omitempty"`
}

type OutputModelElementDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type OutputModelElementsDTO struct {
	Elements []OutputModelElementDTO `json:"elements,omitempty"`
}

type PropertiesIdDTO struct {
	Properties []int `json:"properties_id,omitempty"`
}

type PropertyNamesDTO struct {
	Properties []string `json:"property_names,omitempty"`
}

type OutputPropertyDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type OutputPropertiesDTO struct {
	Properties []OutputPropertyDTO `json:"properties,omitempty"`
}

type SearchUnitDTO struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

type SaveUnitDTO struct {
	Text         string `json:"text"`
	ModelID      int    `json:"model_id"`
	PropertiesID []int  `json:"properties_id,omitempty"`
}

type SaveUnitsDTO struct {
	Contexts map[string]string        `json:"contexts,omitempty"`
	Units    []map[string]SaveUnitDTO `json:"units,omitempty"`
}

type OutputUnitDTO struct {
	ModelID      int    `json:"model_id"`
	RegDate      string `json:"reg_date"`
	Text         string `json:"text"`
	PropertiesID []int  `json:"properties_id,omitempty"`
	ContextsID   []int  `json:"contexts_id,omitempty"`
}

type UnitDtoMaps []map[string]OutputUnitDTO

type ContextDTO struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type OutputUnitsDTO struct {
	Units    []map[string]OutputUnitDTO `json:"units,omitempty"`
	Contexts []ContextDTO               `json:"contexts,omitempty"`
}

type UpdateUnitDTO struct {
	Lang         string  `json:"lang"`
	OldText      string  `json:"old_text"`
	NewText      *string `json:"new_text,omitempty"`
	PropertiesID []int   `json:"properties_id,omitempty"`
}

type UpdateUnitsDTO struct {
	Units []UpdateUnitDTO `json:"units,omitempty"`
}

type LayersDTO struct {
	Layers []string `json:"layers,omitempty"`
}
