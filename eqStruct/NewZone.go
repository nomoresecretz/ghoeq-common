package eqStruct

import (
	"github.com/nomoresecretz/ghoeq-common/proto/eqstruct"
	"google.golang.org/protobuf/proto"
)

type NewZone struct {
	Name                 string // 000 Max64
	NameShort            string // 64 Max32
	NameLong             string // 96 Max278
	Type                 uint8
	FogRed               uint8   // 375
	FogGreen             uint8   // 379
	FogBlue              uint8   // 383
	FogMinClip           float32 // 388
	FogMaxClip           float32 // 404
	Gravity              float32 // 420
	TimeType             uint8
	RainChance           uint8  // 425
	RainDuration         uint8  // 429
	SnowChance           uint8  // 433
	SnowDuration         uint8  // 437
	SpecialDates         []byte // 441
	SpecialCodes         []byte // 457 Max16
	TimeZone             int8
	Sky                  uint8
	WaterMusic           int16
	NormalMusicDay       int16
	NormalMusicNight     int16
	ZoneExpMult          float32
	SafeY                float32
	SafeX                float32
	SafeZ                float32
	MaxZ                 float32
	Underworld           float32
	ClipMin              float32
	ClipMax              float32
	ForageNovice         uint32
	ForageMedium         uint32
	ForageAdv            uint32
	FishingNovice        uint32
	FishingMedium        uint32
	FishingAdv           uint32
	SkyLock              uint32
	GraveyardTime        uint16
	ScriptPeriodicHour   uint32
	ScriptPeriodicMinute uint32
	ScriptPeriodicFast   uint32
	ScriptPlayerDead     uint32
	ScriptNPCDead        uint32
	ScriptPlayerEntering uint32

	bPointer int
}

func (p *NewZone) EQType() EQType   { return EQT_NewZone }
func (p *NewZone) bp() *int         { return &p.bPointer }
func (p *NewZone) SetPointer(i int) { p.bPointer = i }

func (p *NewZone) Unmarshal(b []byte) (int, error) {
	p.bPointer = 0

	if err := EQRead(b, p, &p.Name, 64); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.NameShort, 32); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.NameLong, 278); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Type, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.FogRed, 0); err != nil {
		return 0, err
	}

	p.bPointer = 379
	if err := EQRead(b, p, &p.FogGreen, 0); err != nil {
		return 0, err
	}

	p.bPointer = 383
	if err := EQRead(b, p, &p.FogBlue, 0); err != nil {
		return 0, err
	}

	p.bPointer = 388
	if err := EQRead(b, p, &p.FogMinClip, 0); err != nil {
		return 0, err
	}

	p.bPointer = 404
	if err := EQRead(b, p, &p.FogMaxClip, 0); err != nil {
		return 0, err
	}

	p.bPointer = 420
	if err := EQRead(b, p, &p.Gravity, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.TimeType, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.RainChance, 0); err != nil {
		return 0, err
	}

	p.bPointer = 429
	if err := EQRead(b, p, &p.RainDuration, 0); err != nil {
		return 0, err
	}

	p.bPointer = 433
	if err := EQRead(b, p, &p.SnowChance, 0); err != nil {
		return 0, err
	}

	p.bPointer = 437
	if err := EQRead(b, p, &p.SnowDuration, 0); err != nil {
		return 0, err
	}

	p.bPointer = 441
	if err := EQRead(b, p, &p.SpecialDates, 4); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SpecialCodes, 4); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.TimeZone, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Sky, 0); err != nil {
		return 0, err
	}

	p.bPointer = 476
	if err := EQRead(b, p, &p.WaterMusic, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.NormalMusicDay, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.NormalMusicNight, 0); err != nil {
		return 0, err
	}

	p.bPointer = 484
	if err := EQRead(b, p, &p.ZoneExpMult, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SafeY, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SafeX, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SafeZ, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.MaxZ, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.Underworld, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ClipMin, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ClipMax, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ForageNovice, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ForageMedium, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ForageAdv, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.FishingNovice, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.FishingMedium, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.FishingAdv, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.SkyLock, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.GraveyardTime, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptPeriodicHour, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptPeriodicMinute, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptPeriodicFast, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptPlayerDead, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptNPCDead, 0); err != nil {
		return 0, err
	}

	if err := EQRead(b, p, &p.ScriptPlayerEntering, 0); err != nil {
		return 0, err
	}

	return p.bPointer, nil
}

func (p *NewZone) Proto() *eqstruct.NewZone {
	return &eqstruct.NewZone{
		Name:                 p.Name,
		NameShort:            p.NameShort,
		NameLong:             p.NameLong,
		Type:                 uint32(p.Type),
		FogRed:               uint32(p.FogRed),
		FogGreen:             uint32(p.FogGreen),
		FogBlue:              uint32(p.FogBlue),
		FogMinClip:           p.FogMinClip,
		FogMaxClip:           p.FogMaxClip,
		Gravity:              p.Gravity,
		TimeType:             uint32(p.TimeType),
		RainChance:           uint32(p.RainChance),
		RainDuration:         uint32(p.RainDuration),
		SnowChance:           uint32(p.SnowChance),
		SnowDuration:         uint32(p.SnowDuration),
		SpecialDates:         p.SpecialDates,
		SpecialCodes:         p.SpecialCodes,
		TimeZone:             uint32(p.TimeZone),
		Sky:                  uint32(p.Sky),
		WaterMusic:           int32(p.WaterMusic),
		MusicNormalDay:       int32(p.NormalMusicDay),
		MusicNormalNight:     int32(p.NormalMusicNight),
		ZoneExpMultiplier:    p.ZoneExpMult,
		SafeX:                p.SafeX,
		SafeY:                p.SafeY,
		SafeZ:                p.SafeZ,
		MaxZ:                 p.MaxZ,
		Underworld:           p.Underworld,
		ClipMin:              p.ClipMin,
		ClipMax:              p.ClipMax,
		ForageNovice:         p.ForageNovice,
		ForageMedium:         p.ForageMedium,
		ForageAdvance:        p.ForageAdv,
		FishingNovice:        p.FishingNovice,
		FishingMedium:        p.FishingMedium,
		FishingAdvance:       p.FishingAdv,
		SkyLock:              p.SkyLock,
		GraveyardTime:        uint32(p.GraveyardTime),
		ScriptPeriodicHour:   p.ScriptPeriodicHour,
		ScriptPeriodicMin:    p.ScriptPeriodicMinute,
		ScriptPeriodicFast:   p.ScriptPeriodicFast,
		ScriptPlayerDead:     p.ScriptPlayerDead,
		ScriptNpcDead:        p.ScriptNPCDead,
		ScriptPlayerEntering: p.ScriptPlayerEntering,
	}
}

func (p *NewZone) ProtoMess() proto.Message {
	return p.Proto()
}
