package zeewImg

type BienvenidaEstilo string

const (
	Classic BienvenidaEstilo = "classic"
	Anime   BienvenidaEstilo = "anime"
	Ani     BienvenidaEstilo = "ani"
	Simple  BienvenidaEstilo = "simple"
)

type Bienvenida struct {
	Token       string           `json:"token"`
	Tit         string           `json:"tit"`
	ColorTit    string           `json:"colortit"`
	ColorDesc   string           `json:"colordesc"`
	Desc        string           `json:"desc"`
	Avatar      string           `json:"avatar"`
	Fondo       string           `json:"fondo"`
	Estilo      BienvenidaEstilo `json:"estilo"`
	ColorCircle string           `json:"colorCircle"`
	Font        string           `json:"font"`
}

func NewBienvenida(token string) *Bienvenida {
	return &Bienvenida{
		Token: token,
	}
}

// Token establece el valor del token y retorna la instancia actual.
func (b *Bienvenida) SetToken(token string) *Bienvenida {
	b.Token = token
	return b
}

// Titulo establece el valor del título y retorna la instancia actual.
func (b *Bienvenida) Titulo(tit string) *Bienvenida {
	b.Tit = tit
	return b
}

// ColorTit establece el color del título y retorna la instancia actual.
func (b *Bienvenida) SetColorTit(colortit string) *Bienvenida {
	b.ColorTit = colortit
	return b
}

// ColorDesc establece el color de la descripción y retorna la instancia actual.
func (b *Bienvenida) SetColorDesc(colordesc string) *Bienvenida {
	b.ColorDesc = colordesc
	return b
}

// Descripcion establece la descripción y retorna la instancia actual.
func (b *Bienvenida) Descripcion(desc string) *Bienvenida {
	b.Desc = desc
	return b
}

// Avatar establece el avatar y retorna la instancia actual.
func (b *Bienvenida) SetAvatar(avatar string) *Bienvenida {
	b.Avatar = avatar
	return b
}

func (b *Bienvenida) SetFondo(fondo string) *Bienvenida {
	b.Fondo = fondo
	return b
}

func (b *Bienvenida) SetEstilo(estilo BienvenidaEstilo) *Bienvenida {
	b.Estilo = estilo
	return b
}

// ColorCirculo establece el color del círculo y retorna la instancia actual.
func (b *Bienvenida) ColorCirculo(colorCircle string) *Bienvenida {
	b.ColorCircle = colorCircle
	return b
}

// Font establece la fuente y retorna la instancia actual.
func (b *Bienvenida) SetFont(font string) *Bienvenida {
	b.Font = font
	return b
}
