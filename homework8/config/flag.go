package config

type StringFlag struct {
    Setb   bool
    Value string
}

func (sf *StringFlag) Set(x string) error {
    sf.Value = x
    sf.Setb = true
    return nil
}

func (sf *StringFlag) String() string {
    return sf.Value
}

