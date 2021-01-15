package epub

//Ncx OPS/toc.ncx
type Ncx struct {
	NavMap []NavPoint `xml:"navMap>navPoint" json:"navMap"`
}

//NavPoint nav point
type NavPoint struct {
	Text      string     `xml:"navLabel>text" json:"text"`
	Content   Content    `xml:"content" json:"content"`
	Points    []NavPoint `xml:"navPoint" json:"points"`
	PlayOrder string     `xml:"playOrder,attr" json:"playOrder"`
}

//Content nav-point content
type Content struct {
	Src string `xml:"src,attr" json:"src"`
}
