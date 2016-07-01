// Code generated by protoc-gen-go.
// source: metadata.proto
// DO NOT EDIT!

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Album_Type int32

const (
	Album_ALBUM       Album_Type = 1
	Album_SINGLE      Album_Type = 2
	Album_COMPILATION Album_Type = 3
)

var Album_Type_name = map[int32]string{
	1: "ALBUM",
	2: "SINGLE",
	3: "COMPILATION",
}
var Album_Type_value = map[string]int32{
	"ALBUM":       1,
	"SINGLE":      2,
	"COMPILATION": 3,
}

func (x Album_Type) Enum() *Album_Type {
	p := new(Album_Type)
	*p = x
	return p
}
func (x Album_Type) String() string {
	return proto.EnumName(Album_Type_name, int32(x))
}
func (x *Album_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Album_Type_value, data, "Album_Type")
	if err != nil {
		return err
	}
	*x = Album_Type(value)
	return nil
}
func (Album_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5, 0} }

type Image_Size int32

const (
	Image_DEFAULT Image_Size = 0
	Image_SMALL   Image_Size = 1
	Image_LARGE   Image_Size = 2
	Image_XLARGE  Image_Size = 3
)

var Image_Size_name = map[int32]string{
	0: "DEFAULT",
	1: "SMALL",
	2: "LARGE",
	3: "XLARGE",
}
var Image_Size_value = map[string]int32{
	"DEFAULT": 0,
	"SMALL":   1,
	"LARGE":   2,
	"XLARGE":  3,
}

func (x Image_Size) Enum() *Image_Size {
	p := new(Image_Size)
	*p = x
	return p
}
func (x Image_Size) String() string {
	return proto.EnumName(Image_Size_name, int32(x))
}
func (x *Image_Size) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Image_Size_value, data, "Image_Size")
	if err != nil {
		return err
	}
	*x = Image_Size(value)
	return nil
}
func (Image_Size) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{7, 0} }

type Copyright_Type int32

const (
	Copyright_P Copyright_Type = 0
	Copyright_C Copyright_Type = 1
)

var Copyright_Type_name = map[int32]string{
	0: "P",
	1: "C",
}
var Copyright_Type_value = map[string]int32{
	"P": 0,
	"C": 1,
}

func (x Copyright_Type) Enum() *Copyright_Type {
	p := new(Copyright_Type)
	*p = x
	return p
}
func (x Copyright_Type) String() string {
	return proto.EnumName(Copyright_Type_name, int32(x))
}
func (x *Copyright_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Copyright_Type_value, data, "Copyright_Type")
	if err != nil {
		return err
	}
	*x = Copyright_Type(value)
	return nil
}
func (Copyright_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{11, 0} }

type Restriction_Type int32

const (
	Restriction_STREAMING Restriction_Type = 0
)

var Restriction_Type_name = map[int32]string{
	0: "STREAMING",
}
var Restriction_Type_value = map[string]int32{
	"STREAMING": 0,
}

func (x Restriction_Type) Enum() *Restriction_Type {
	p := new(Restriction_Type)
	*p = x
	return p
}
func (x Restriction_Type) String() string {
	return proto.EnumName(Restriction_Type_name, int32(x))
}
func (x *Restriction_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Restriction_Type_value, data, "Restriction_Type")
	if err != nil {
		return err
	}
	*x = Restriction_Type(value)
	return nil
}
func (Restriction_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{12, 0} }

type AudioFile_Format int32

const (
	AudioFile_OGG_VORBIS_96  AudioFile_Format = 0
	AudioFile_OGG_VORBIS_160 AudioFile_Format = 1
	AudioFile_OGG_VORBIS_320 AudioFile_Format = 2
	AudioFile_MP3_256        AudioFile_Format = 3
	AudioFile_MP3_320        AudioFile_Format = 4
	AudioFile_MP3_160        AudioFile_Format = 5
	AudioFile_MP3_96         AudioFile_Format = 6
	AudioFile_MP3_160_ENC    AudioFile_Format = 7
	AudioFile_OTHER2         AudioFile_Format = 8
	AudioFile_OTHER3         AudioFile_Format = 9
	AudioFile_AAC_160        AudioFile_Format = 10
	AudioFile_AAC_320        AudioFile_Format = 11
)

var AudioFile_Format_name = map[int32]string{
	0:  "OGG_VORBIS_96",
	1:  "OGG_VORBIS_160",
	2:  "OGG_VORBIS_320",
	3:  "MP3_256",
	4:  "MP3_320",
	5:  "MP3_160",
	6:  "MP3_96",
	7:  "MP3_160_ENC",
	8:  "OTHER2",
	9:  "OTHER3",
	10: "AAC_160",
	11: "AAC_320",
}
var AudioFile_Format_value = map[string]int32{
	"OGG_VORBIS_96":  0,
	"OGG_VORBIS_160": 1,
	"OGG_VORBIS_320": 2,
	"MP3_256":        3,
	"MP3_320":        4,
	"MP3_160":        5,
	"MP3_96":         6,
	"MP3_160_ENC":    7,
	"OTHER2":         8,
	"OTHER3":         9,
	"AAC_160":        10,
	"AAC_320":        11,
}

func (x AudioFile_Format) Enum() *AudioFile_Format {
	p := new(AudioFile_Format)
	*p = x
	return p
}
func (x AudioFile_Format) String() string {
	return proto.EnumName(AudioFile_Format_name, int32(x))
}
func (x *AudioFile_Format) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AudioFile_Format_value, data, "AudioFile_Format")
	if err != nil {
		return err
	}
	*x = AudioFile_Format(value)
	return nil
}
func (AudioFile_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{15, 0} }

type TopTracks struct {
	Country          *string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	Track            []*Track `protobuf:"bytes,2,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TopTracks) Reset()                    { *m = TopTracks{} }
func (m *TopTracks) String() string            { return proto.CompactTextString(m) }
func (*TopTracks) ProtoMessage()               {}
func (*TopTracks) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *TopTracks) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *TopTracks) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type ActivityPeriod struct {
	StartYear        *int32 `protobuf:"zigzag32,1,opt,name=start_year" json:"start_year,omitempty"`
	EndYear          *int32 `protobuf:"zigzag32,2,opt,name=end_year" json:"end_year,omitempty"`
	Decade           *int32 `protobuf:"zigzag32,3,opt,name=decade" json:"decade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActivityPeriod) Reset()                    { *m = ActivityPeriod{} }
func (m *ActivityPeriod) String() string            { return proto.CompactTextString(m) }
func (*ActivityPeriod) ProtoMessage()               {}
func (*ActivityPeriod) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *ActivityPeriod) GetStartYear() int32 {
	if m != nil && m.StartYear != nil {
		return *m.StartYear
	}
	return 0
}

func (m *ActivityPeriod) GetEndYear() int32 {
	if m != nil && m.EndYear != nil {
		return *m.EndYear
	}
	return 0
}

func (m *ActivityPeriod) GetDecade() int32 {
	if m != nil && m.Decade != nil {
		return *m.Decade
	}
	return 0
}

type Artist struct {
	Gid                  []byte            `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name                 *string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Popularity           *int32            `protobuf:"zigzag32,3,opt,name=popularity" json:"popularity,omitempty"`
	TopTrack             []*TopTracks      `protobuf:"bytes,4,rep,name=top_track" json:"top_track,omitempty"`
	AlbumGroup           []*AlbumGroup     `protobuf:"bytes,5,rep,name=album_group" json:"album_group,omitempty"`
	SingleGroup          []*AlbumGroup     `protobuf:"bytes,6,rep,name=single_group" json:"single_group,omitempty"`
	CompilationGroup     []*AlbumGroup     `protobuf:"bytes,7,rep,name=compilation_group" json:"compilation_group,omitempty"`
	AppearsOnGroup       []*AlbumGroup     `protobuf:"bytes,8,rep,name=appears_on_group" json:"appears_on_group,omitempty"`
	Genre                []string          `protobuf:"bytes,9,rep,name=genre" json:"genre,omitempty"`
	ExternalId           []*ExternalId     `protobuf:"bytes,10,rep,name=external_id" json:"external_id,omitempty"`
	Portrait             []*Image          `protobuf:"bytes,11,rep,name=portrait" json:"portrait,omitempty"`
	Biography            []*Biography      `protobuf:"bytes,12,rep,name=biography" json:"biography,omitempty"`
	ActivityPeriod       []*ActivityPeriod `protobuf:"bytes,13,rep,name=activity_period" json:"activity_period,omitempty"`
	Restriction          []*Restriction    `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related              []*Artist         `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	IsPortraitAlbumCover *bool             `protobuf:"varint,16,opt,name=is_portrait_album_cover" json:"is_portrait_album_cover,omitempty"`
	PortraitGroup        *ImageGroup       `protobuf:"bytes,17,opt,name=portrait_group" json:"portrait_group,omitempty"`
	XXX_unrecognized     []byte            `json:"-"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *Artist) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Artist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artist) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Artist) GetTopTrack() []*TopTracks {
	if m != nil {
		return m.TopTrack
	}
	return nil
}

func (m *Artist) GetAlbumGroup() []*AlbumGroup {
	if m != nil {
		return m.AlbumGroup
	}
	return nil
}

func (m *Artist) GetSingleGroup() []*AlbumGroup {
	if m != nil {
		return m.SingleGroup
	}
	return nil
}

func (m *Artist) GetCompilationGroup() []*AlbumGroup {
	if m != nil {
		return m.CompilationGroup
	}
	return nil
}

func (m *Artist) GetAppearsOnGroup() []*AlbumGroup {
	if m != nil {
		return m.AppearsOnGroup
	}
	return nil
}

func (m *Artist) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Artist) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Artist) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Artist) GetBiography() []*Biography {
	if m != nil {
		return m.Biography
	}
	return nil
}

func (m *Artist) GetActivityPeriod() []*ActivityPeriod {
	if m != nil {
		return m.ActivityPeriod
	}
	return nil
}

func (m *Artist) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Artist) GetRelated() []*Artist {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Artist) GetIsPortraitAlbumCover() bool {
	if m != nil && m.IsPortraitAlbumCover != nil {
		return *m.IsPortraitAlbumCover
	}
	return false
}

func (m *Artist) GetPortraitGroup() *ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type AlbumGroup struct {
	Album            []*Album `protobuf:"bytes,1,rep,name=album" json:"album,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AlbumGroup) Reset()                    { *m = AlbumGroup{} }
func (m *AlbumGroup) String() string            { return proto.CompactTextString(m) }
func (*AlbumGroup) ProtoMessage()               {}
func (*AlbumGroup) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *AlbumGroup) GetAlbum() []*Album {
	if m != nil {
		return m.Album
	}
	return nil
}

type Date struct {
	Year             *int32 `protobuf:"zigzag32,1,opt,name=year" json:"year,omitempty"`
	Month            *int32 `protobuf:"zigzag32,2,opt,name=month" json:"month,omitempty"`
	Day              *int32 `protobuf:"zigzag32,3,opt,name=day" json:"day,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *Date) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

type Album struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,3,rep,name=artist" json:"artist,omitempty"`
	Typ              *Album_Type    `protobuf:"varint,4,opt,name=typ,enum=Spotify.Album_Type" json:"typ,omitempty"`
	Label            *string        `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Date             *Date          `protobuf:"bytes,6,opt,name=date" json:"date,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,7,opt,name=popularity" json:"popularity,omitempty"`
	Genre            []string       `protobuf:"bytes,8,rep,name=genre" json:"genre,omitempty"`
	Cover            []*Image       `protobuf:"bytes,9,rep,name=cover" json:"cover,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id" json:"external_id,omitempty"`
	Disc             []*Disc        `protobuf:"bytes,11,rep,name=disc" json:"disc,omitempty"`
	Review           []string       `protobuf:"bytes,12,rep,name=review" json:"review,omitempty"`
	Copyright        []*Copyright   `protobuf:"bytes,13,rep,name=copyright" json:"copyright,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related          []*Album       `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,16,rep,name=sale_period" json:"sale_period,omitempty"`
	CoverGroup       *ImageGroup    `protobuf:"bytes,17,opt,name=cover_group" json:"cover_group,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Album) Reset()                    { *m = Album{} }
func (m *Album) String() string            { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()               {}
func (*Album) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *Album) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Album) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Album) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Album) GetTyp() Album_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Album_ALBUM
}

func (m *Album) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Album) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Album) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Album) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Album) GetCover() []*Image {
	if m != nil {
		return m.Cover
	}
	return nil
}

func (m *Album) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Album) GetDisc() []*Disc {
	if m != nil {
		return m.Disc
	}
	return nil
}

func (m *Album) GetReview() []string {
	if m != nil {
		return m.Review
	}
	return nil
}

func (m *Album) GetCopyright() []*Copyright {
	if m != nil {
		return m.Copyright
	}
	return nil
}

func (m *Album) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Album) GetRelated() []*Album {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Album) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Album) GetCoverGroup() *ImageGroup {
	if m != nil {
		return m.CoverGroup
	}
	return nil
}

type Track struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Album            *Album         `protobuf:"bytes,3,opt,name=album" json:"album,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,4,rep,name=artist" json:"artist,omitempty"`
	Number           *int32         `protobuf:"zigzag32,5,opt,name=number" json:"number,omitempty"`
	DiscNumber       *int32         `protobuf:"zigzag32,6,opt,name=disc_number" json:"disc_number,omitempty"`
	Duration         *int32         `protobuf:"zigzag32,7,opt,name=duration" json:"duration,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,8,opt,name=popularity" json:"popularity,omitempty"`
	Explicit         *bool          `protobuf:"varint,9,opt,name=explicit" json:"explicit,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id" json:"external_id,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,11,rep,name=restriction" json:"restriction,omitempty"`
	File             []*AudioFile   `protobuf:"bytes,12,rep,name=file" json:"file,omitempty"`
	Alternative      []*Track       `protobuf:"bytes,13,rep,name=alternative" json:"alternative,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,14,rep,name=sale_period" json:"sale_period,omitempty"`
	Preview          []*AudioFile   `protobuf:"bytes,15,rep,name=preview" json:"preview,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}
func (*Track) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

func (m *Track) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Track) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Track) GetAlbum() *Album {
	if m != nil {
		return m.Album
	}
	return nil
}

func (m *Track) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Track) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Track) GetDiscNumber() int32 {
	if m != nil && m.DiscNumber != nil {
		return *m.DiscNumber
	}
	return 0
}

func (m *Track) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Track) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Track) GetExplicit() bool {
	if m != nil && m.Explicit != nil {
		return *m.Explicit
	}
	return false
}

func (m *Track) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Track) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Track) GetFile() []*AudioFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Track) GetAlternative() []*Track {
	if m != nil {
		return m.Alternative
	}
	return nil
}

func (m *Track) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Track) GetPreview() []*AudioFile {
	if m != nil {
		return m.Preview
	}
	return nil
}

type Image struct {
	FileId           []byte      `protobuf:"bytes,1,opt,name=file_id" json:"file_id,omitempty"`
	Size             *Image_Size `protobuf:"varint,2,opt,name=size,enum=Spotify.Image_Size" json:"size,omitempty"`
	Width            *int32      `protobuf:"zigzag32,3,opt,name=width" json:"width,omitempty"`
	Height           *int32      `protobuf:"zigzag32,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *Image) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *Image) GetSize() Image_Size {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return Image_DEFAULT
}

func (m *Image) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type ImageGroup struct {
	Image            []*Image `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImageGroup) Reset()                    { *m = ImageGroup{} }
func (m *ImageGroup) String() string            { return proto.CompactTextString(m) }
func (*ImageGroup) ProtoMessage()               {}
func (*ImageGroup) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ImageGroup) GetImage() []*Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Biography struct {
	Text             *string       `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Portrait         []*Image      `protobuf:"bytes,2,rep,name=portrait" json:"portrait,omitempty"`
	PortraitGroup    []*ImageGroup `protobuf:"bytes,3,rep,name=portrait_group" json:"portrait_group,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Biography) Reset()                    { *m = Biography{} }
func (m *Biography) String() string            { return proto.CompactTextString(m) }
func (*Biography) ProtoMessage()               {}
func (*Biography) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

func (m *Biography) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Biography) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Biography) GetPortraitGroup() []*ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type Disc struct {
	Number           *int32   `protobuf:"zigzag32,1,opt,name=number" json:"number,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Track            []*Track `protobuf:"bytes,3,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Disc) Reset()                    { *m = Disc{} }
func (m *Disc) String() string            { return proto.CompactTextString(m) }
func (*Disc) ProtoMessage()               {}
func (*Disc) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

func (m *Disc) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Disc) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Disc) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type Copyright struct {
	Typ              *Copyright_Type `protobuf:"varint,1,opt,name=typ,enum=Spotify.Copyright_Type" json:"typ,omitempty"`
	Text             *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Copyright) Reset()                    { *m = Copyright{} }
func (m *Copyright) String() string            { return proto.CompactTextString(m) }
func (*Copyright) ProtoMessage()               {}
func (*Copyright) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

func (m *Copyright) GetTyp() Copyright_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Copyright_P
}

func (m *Copyright) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type Restriction struct {
	CountriesAllowed   *string           `protobuf:"bytes,2,opt,name=countries_allowed" json:"countries_allowed,omitempty"`
	CountriesForbidden *string           `protobuf:"bytes,3,opt,name=countries_forbidden" json:"countries_forbidden,omitempty"`
	Typ                *Restriction_Type `protobuf:"varint,4,opt,name=typ,enum=Spotify.Restriction_Type" json:"typ,omitempty"`
	CatalogueStr       []string          `protobuf:"bytes,5,rep,name=catalogue_str" json:"catalogue_str,omitempty"`
	XXX_unrecognized   []byte            `json:"-"`
}

func (m *Restriction) Reset()                    { *m = Restriction{} }
func (m *Restriction) String() string            { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()               {}
func (*Restriction) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

func (m *Restriction) GetCountriesAllowed() string {
	if m != nil && m.CountriesAllowed != nil {
		return *m.CountriesAllowed
	}
	return ""
}

func (m *Restriction) GetCountriesForbidden() string {
	if m != nil && m.CountriesForbidden != nil {
		return *m.CountriesForbidden
	}
	return ""
}

func (m *Restriction) GetTyp() Restriction_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Restriction_STREAMING
}

func (m *Restriction) GetCatalogueStr() []string {
	if m != nil {
		return m.CatalogueStr
	}
	return nil
}

type SalePeriod struct {
	Restriction      []*Restriction `protobuf:"bytes,1,rep,name=restriction" json:"restriction,omitempty"`
	Start            *Date          `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End              *Date          `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SalePeriod) Reset()                    { *m = SalePeriod{} }
func (m *SalePeriod) String() string            { return proto.CompactTextString(m) }
func (*SalePeriod) ProtoMessage()               {}
func (*SalePeriod) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *SalePeriod) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *SalePeriod) GetStart() *Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SalePeriod) GetEnd() *Date {
	if m != nil {
		return m.End
	}
	return nil
}

type ExternalId struct {
	Typ              *string `protobuf:"bytes,1,opt,name=typ" json:"typ,omitempty"`
	Id               *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalId) Reset()                    { *m = ExternalId{} }
func (m *ExternalId) String() string            { return proto.CompactTextString(m) }
func (*ExternalId) ProtoMessage()               {}
func (*ExternalId) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{14} }

func (m *ExternalId) GetTyp() string {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return ""
}

func (m *ExternalId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type AudioFile struct {
	FileId           []byte            `protobuf:"bytes,1,opt,name=file_id" json:"file_id,omitempty"`
	Format           *AudioFile_Format `protobuf:"varint,2,opt,name=format,enum=Spotify.AudioFile_Format" json:"format,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *AudioFile) Reset()                    { *m = AudioFile{} }
func (m *AudioFile) String() string            { return proto.CompactTextString(m) }
func (*AudioFile) ProtoMessage()               {}
func (*AudioFile) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{15} }

func (m *AudioFile) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *AudioFile) GetFormat() AudioFile_Format {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return AudioFile_OGG_VORBIS_96
}

func init() {
	proto.RegisterType((*TopTracks)(nil), "Spotify.TopTracks")
	proto.RegisterType((*ActivityPeriod)(nil), "Spotify.ActivityPeriod")
	proto.RegisterType((*Artist)(nil), "Spotify.Artist")
	proto.RegisterType((*AlbumGroup)(nil), "Spotify.AlbumGroup")
	proto.RegisterType((*Date)(nil), "Spotify.Date")
	proto.RegisterType((*Album)(nil), "Spotify.Album")
	proto.RegisterType((*Track)(nil), "Spotify.Track")
	proto.RegisterType((*Image)(nil), "Spotify.Image")
	proto.RegisterType((*ImageGroup)(nil), "Spotify.ImageGroup")
	proto.RegisterType((*Biography)(nil), "Spotify.Biography")
	proto.RegisterType((*Disc)(nil), "Spotify.Disc")
	proto.RegisterType((*Copyright)(nil), "Spotify.Copyright")
	proto.RegisterType((*Restriction)(nil), "Spotify.Restriction")
	proto.RegisterType((*SalePeriod)(nil), "Spotify.SalePeriod")
	proto.RegisterType((*ExternalId)(nil), "Spotify.ExternalId")
	proto.RegisterType((*AudioFile)(nil), "Spotify.AudioFile")
	proto.RegisterEnum("Spotify.Album_Type", Album_Type_name, Album_Type_value)
	proto.RegisterEnum("Spotify.Image_Size", Image_Size_name, Image_Size_value)
	proto.RegisterEnum("Spotify.Copyright_Type", Copyright_Type_name, Copyright_Type_value)
	proto.RegisterEnum("Spotify.Restriction_Type", Restriction_Type_name, Restriction_Type_value)
	proto.RegisterEnum("Spotify.AudioFile_Format", AudioFile_Format_name, AudioFile_Format_value)
}

var fileDescriptor6 = []byte{
	// 1207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x4e, 0xe3, 0x46,
	0x14, 0xc7, 0x37, 0xb1, 0xf3, 0xe1, 0x63, 0x12, 0x8c, 0xd9, 0x15, 0xa1, 0x1f, 0x82, 0x7a, 0xbb,
	0x15, 0x68, 0xd5, 0x88, 0x82, 0x16, 0xa9, 0xea, 0x55, 0xc8, 0x06, 0x1a, 0x29, 0x7c, 0x88, 0x84,
	0xaa, 0x77, 0x96, 0xb1, 0x87, 0x30, 0xaa, 0x13, 0xbb, 0xf6, 0x04, 0x36, 0xbd, 0xed, 0x13, 0xf4,
	0xba, 0x57, 0xbd, 0xed, 0x1b, 0xf4, 0x71, 0xfa, 0x26, 0x3d, 0x33, 0x9e, 0xc4, 0x8e, 0x13, 0xe8,
	0xee, 0x55, 0xec, 0x39, 0x67, 0x26, 0xe7, 0xe3, 0x77, 0xfe, 0x63, 0xa8, 0x8f, 0x08, 0x73, 0x3c,
	0x87, 0x39, 0xcd, 0x30, 0x0a, 0x58, 0x60, 0x56, 0xfa, 0x61, 0xc0, 0xe8, 0xdd, 0xd4, 0xfa, 0x01,
	0xb4, 0x41, 0x10, 0x0e, 0x22, 0xc7, 0xfd, 0x25, 0x36, 0xd7, 0xa1, 0xe2, 0x06, 0x93, 0x31, 0x8b,
	0xa6, 0x8d, 0xc2, 0x6e, 0x61, 0x4f, 0x33, 0xbf, 0x84, 0x12, 0xe3, 0xa6, 0x46, 0x71, 0x57, 0xd9,
	0xd3, 0x0f, 0xeb, 0x4d, 0xb9, 0xad, 0x29, 0x36, 0x58, 0xa7, 0x50, 0x6f, 0xb9, 0x8c, 0x3e, 0x50,
	0x36, 0xbd, 0x22, 0x11, 0x0d, 0x3c, 0xd3, 0x04, 0x88, 0x99, 0x13, 0x31, 0x7b, 0x4a, 0x9c, 0x48,
	0x1c, 0xb2, 0x61, 0x1a, 0x50, 0x25, 0x63, 0x2f, 0x59, 0x29, 0x8a, 0x95, 0x3a, 0x94, 0x3d, 0xe2,
	0x3a, 0x1e, 0x69, 0x28, 0xfc, 0xdd, 0xfa, 0x57, 0x85, 0x72, 0x2b, 0x62, 0x34, 0x66, 0xa6, 0x0e,
	0xca, 0x90, 0x7a, 0x62, 0xe7, 0x9a, 0xb9, 0x06, 0xea, 0xd8, 0x19, 0x11, 0xb1, 0x4b, 0xe3, 0x67,
	0x87, 0x41, 0x38, 0xf1, 0x9d, 0x08, 0xff, 0x2f, 0xd9, 0x69, 0xbe, 0x01, 0x8d, 0x05, 0xa1, 0x9d,
	0x04, 0xa9, 0x8a, 0x20, 0xcd, 0x34, 0xc8, 0x79, 0x62, 0x7b, 0xa0, 0x3b, 0xfe, 0xed, 0x64, 0x64,
	0x0f, 0xa3, 0x60, 0x12, 0x36, 0x4a, 0xc2, 0x71, 0x73, 0xee, 0xd8, 0xe2, 0xb6, 0x33, 0x6e, 0x32,
	0xf7, 0x61, 0x2d, 0xa6, 0xe3, 0xa1, 0x4f, 0xa4, 0x6b, 0xf9, 0x69, 0xd7, 0x26, 0x6c, 0xb8, 0xc1,
	0x28, 0xa4, 0xbe, 0xc3, 0x68, 0x30, 0x96, 0xfe, 0x95, 0xa7, 0xfd, 0xbf, 0x05, 0xc3, 0x09, 0x43,
	0xac, 0x42, 0x6c, 0xcf, 0xdd, 0xab, 0x4f, 0xbb, 0xd7, 0xa0, 0x34, 0x24, 0xe3, 0x88, 0x34, 0x34,
	0xf4, 0xd1, 0x78, 0x0a, 0xe4, 0x03, 0x23, 0xd1, 0xd8, 0xf1, 0x6d, 0x2c, 0x10, 0xe4, 0x36, 0x76,
	0xa4, 0xad, 0xeb, 0x99, 0xbb, 0x50, 0x0d, 0x83, 0x08, 0x4b, 0x42, 0x59, 0x43, 0xcf, 0xf5, 0xad,
	0x3b, 0x72, 0x86, 0x84, 0x57, 0xed, 0x96, 0x06, 0xc3, 0xc8, 0x09, 0xef, 0xa7, 0x8d, 0xb5, 0x5c,
	0xd5, 0x4e, 0x66, 0x16, 0xf3, 0x00, 0xd6, 0x1d, 0xd9, 0x5e, 0x3b, 0x14, 0xfd, 0x6d, 0xd4, 0x84,
	0xf3, 0x56, 0x1a, 0xef, 0x62, 0xfb, 0xf7, 0x41, 0x8f, 0x48, 0xcc, 0x22, 0xea, 0xf2, 0x92, 0x34,
	0xea, 0xc2, 0xfb, 0xe5, 0xdc, 0xfb, 0x3a, 0xb5, 0x61, 0x94, 0x95, 0x88, 0x60, 0xe9, 0x88, 0xd7,
	0x58, 0x17, 0x6e, 0xeb, 0xe9, 0xa1, 0x09, 0x0a, 0x3b, 0xb0, 0x45, 0x63, 0x7b, 0x96, 0x8a, 0x9d,
	0x34, 0xd0, 0x0d, 0x1e, 0x48, 0xd4, 0x30, 0xb0, 0xf9, 0x55, 0xf3, 0x2d, 0xd4, 0xe7, 0xd6, 0xa4,
	0x9c, 0x1b, 0xb8, 0x9e, 0xad, 0x8a, 0x48, 0x57, 0x94, 0xd3, 0x7a, 0x0b, 0x90, 0x29, 0x2e, 0x82,
	0x2d, 0xce, 0x43, 0xd0, 0x16, 0x0b, 0x24, 0x7c, 0xac, 0x03, 0x50, 0xdf, 0x63, 0x68, 0x1c, 0xc0,
	0x0c, 0xc8, 0xd8, 0x91, 0x51, 0x30, 0x66, 0xf7, 0x92, 0x62, 0x44, 0xd5, 0x73, 0x24, 0x88, 0xd6,
	0x5f, 0x2a, 0x94, 0xc4, 0xde, 0xe7, 0x08, 0xde, 0x81, 0xb2, 0x23, 0x72, 0xc3, 0x4d, 0x2b, 0x53,
	0xde, 0x05, 0x85, 0x4d, 0x43, 0x04, 0xb9, 0xb0, 0x57, 0xcf, 0x53, 0xd1, 0x1c, 0x4c, 0x43, 0xc2,
	0x63, 0xf0, 0x9d, 0x5b, 0xe2, 0x23, 0xc3, 0xfc, 0xc4, 0xcf, 0x41, 0xc5, 0xa9, 0x26, 0x88, 0x29,
	0x4f, 0xbc, 0x36, 0xdf, 0x21, 0xa2, 0x5f, 0x1c, 0x98, 0xca, 0x2c, 0x87, 0x84, 0xaa, 0xaa, 0xa0,
	0x0a, 0xeb, 0x90, 0x54, 0x54, 0x5b, 0x09, 0xca, 0xc7, 0x43, 0xc7, 0x03, 0xa1, 0xb1, 0x2b, 0x81,
	0xcb, 0x04, 0x82, 0x8b, 0x7c, 0xde, 0x23, 0xf2, 0x40, 0xc9, 0xa3, 0x80, 0x4d, 0xe3, 0xfc, 0xb9,
	0x41, 0x38, 0x8d, 0xe8, 0xf0, 0x9e, 0x49, 0xa4, 0x52, 0xfe, 0xda, 0x33, 0xcb, 0xa7, 0xd0, 0xb4,
	0x93, 0xa7, 0x29, 0xd7, 0x51, 0x9e, 0x49, 0xec, 0xe0, 0x54, 0x4b, 0x8e, 0x8d, 0x5c, 0x26, 0x7d,
	0xb4, 0x49, 0x86, 0xd1, 0x53, 0x94, 0xe4, 0xff, 0x91, 0x6a, 0x82, 0x2a, 0x7a, 0xa2, 0x61, 0xeb,
	0x7b, 0x27, 0x37, 0xe7, 0x46, 0x01, 0x2b, 0x5e, 0xee, 0x77, 0x2f, 0xce, 0x7a, 0x1d, 0xa3, 0x88,
	0x6a, 0xaa, 0xb7, 0x2f, 0xcf, 0xaf, 0xba, 0xbd, 0xd6, 0xa0, 0x7b, 0x79, 0x61, 0x28, 0xd6, 0xdf,
	0x0a, 0x94, 0x84, 0x20, 0x3d, 0xc7, 0xc8, 0x9c, 0x4c, 0x45, 0xfc, 0x71, 0x3e, 0x8f, 0x14, 0x21,
	0x75, 0x35, 0x42, 0x58, 0xeb, 0xf1, 0x64, 0x74, 0x8b, 0x2d, 0x2d, 0x89, 0x86, 0x6f, 0x82, 0xce,
	0x1b, 0x63, 0xcb, 0xc5, 0xf2, 0x4c, 0x92, 0xbd, 0x49, 0x24, 0x74, 0x4b, 0x72, 0xb1, 0xc8, 0x4a,
	0x75, 0x2e, 0xdc, 0x1f, 0x42, 0x9f, 0xba, 0x28, 0x24, 0x9a, 0x98, 0xb8, 0x8f, 0xe7, 0x21, 0xd7,
	0x3b, 0xfd, 0x59, 0x25, 0x50, 0xef, 0xa8, 0x4f, 0x96, 0x84, 0xa8, 0x35, 0xf1, 0x68, 0x70, 0x8a,
	0x16, 0xf3, 0x35, 0x97, 0x6f, 0x71, 0x34, 0xaa, 0x0d, 0x91, 0xc4, 0xe4, 0x2e, 0xa3, 0x7c, 0x87,
	0xeb, 0x4f, 0x77, 0xf8, 0x35, 0x54, 0x42, 0xc9, 0xe3, 0xfa, 0x53, 0xff, 0x69, 0xfd, 0x51, 0x80,
	0x52, 0x32, 0x04, 0x78, 0x2b, 0xf2, 0xf8, 0xec, 0x79, 0xc3, 0xbe, 0x02, 0x35, 0xa6, 0xbf, 0x25,
	0x0d, 0xab, 0xe7, 0xd1, 0x68, 0xf6, 0xd1, 0xc4, 0xc7, 0xec, 0x91, 0x7a, 0x28, 0x15, 0xca, 0xec,
	0xc2, 0xbb, 0x27, 0x82, 0x76, 0x55, 0xa8, 0xc5, 0x3b, 0x50, 0x85, 0x9b, 0x0e, 0x95, 0xf7, 0x9d,
	0xd3, 0xd6, 0x4d, 0x6f, 0x60, 0xbc, 0xe0, 0x18, 0xf5, 0xcf, 0x5b, 0xbd, 0x1e, 0x62, 0x84, 0x8f,
	0xbd, 0xd6, 0xf5, 0x19, 0xa7, 0x08, 0x89, 0xfa, 0x39, 0x79, 0x56, 0xb8, 0x86, 0xa5, 0xf8, 0x71,
	0x52, 0x28, 0x7f, 0x5b, 0xd2, 0x30, 0xe1, 0x63, 0xdd, 0x83, 0x96, 0x4a, 0x39, 0x32, 0xc6, 0xb0,
	0x73, 0xf2, 0x5a, 0xcf, 0xde, 0x10, 0xc5, 0x95, 0x83, 0xbf, 0x2c, 0xad, 0x4a, 0xae, 0x9e, 0x99,
	0x39, 0x68, 0xa3, 0x5a, 0xca, 0x31, 0x97, 0x94, 0x25, 0x7a, 0xb9, 0x04, 0x76, 0x72, 0x4d, 0x2b,
	0x2b, 0xbf, 0x25, 0x6e, 0x40, 0x4b, 0x27, 0xff, 0xeb, 0x44, 0x07, 0x0b, 0xa2, 0xc0, 0x5b, 0xcb,
	0xd2, 0x90, 0x68, 0xe1, 0x2c, 0x29, 0x71, 0xbe, 0xf5, 0x52, 0x4e, 0x63, 0x09, 0x0a, 0x57, 0x58,
	0x4d, 0xfc, 0x69, 0x1b, 0x05, 0xeb, 0xcf, 0x02, 0xe8, 0x59, 0xd8, 0xb6, 0xf9, 0xa5, 0xcd, 0x3f,
	0x71, 0x28, 0x89, 0xf1, 0x4a, 0xf1, 0x83, 0x47, 0x94, 0x8c, 0xa2, 0xd4, 0xd2, 0xcd, 0xd4, 0x74,
	0x17, 0x44, 0xb7, 0xd4, 0xf3, 0xc8, 0x58, 0x74, 0x50, 0x33, 0xbf, 0xc9, 0x2a, 0xf3, 0xf6, 0x2a,
	0x8e, 0x93, 0x98, 0x5e, 0x41, 0xcd, 0xc5, 0xcf, 0x2c, 0x3f, 0x18, 0x4e, 0x88, 0x8d, 0x36, 0xf1,
	0xad, 0xa1, 0x59, 0xaf, 0x64, 0x70, 0x35, 0xd0, 0xfa, 0x83, 0xeb, 0x4e, 0xeb, 0x1c, 0x45, 0xc2,
	0x78, 0x61, 0xfd, 0x0a, 0x90, 0xe1, 0x32, 0x37, 0x33, 0x85, 0x67, 0x66, 0xe6, 0x0b, 0x28, 0x89,
	0xef, 0x2c, 0x11, 0xfa, 0x92, 0xf0, 0x7f, 0x06, 0x0a, 0x7e, 0x71, 0x49, 0x05, 0x59, 0xb4, 0x59,
	0x6f, 0x00, 0x32, 0x63, 0xaa, 0xa7, 0x85, 0xd6, 0x10, 0xb5, 0x22, 0x95, 0xc5, 0xb0, 0x7e, 0x2f,
	0x82, 0x96, 0x0e, 0xe0, 0xd2, 0x08, 0xec, 0x43, 0x19, 0x2b, 0x34, 0x72, 0x98, 0x1c, 0x82, 0xed,
	0xe5, 0x09, 0x6a, 0x9e, 0x0a, 0x07, 0xeb, 0x9f, 0x02, 0x94, 0x93, 0x47, 0x73, 0x03, 0x6a, 0x97,
	0x67, 0x67, 0xf6, 0x4f, 0x97, 0xd7, 0x27, 0xdd, 0xbe, 0xfd, 0xfd, 0x31, 0xb6, 0xc9, 0x84, 0x7a,
	0x66, 0xe9, 0xbb, 0xe3, 0x03, 0xa4, 0x7f, 0x71, 0xed, 0xe8, 0xf0, 0x00, 0xc7, 0x00, 0x27, 0xe5,
	0xfc, 0xea, 0xc8, 0x3e, 0x7c, 0x77, 0x6c, 0x28, 0xb3, 0x17, 0x6e, 0x51, 0x67, 0x2f, 0x7c, 0x6b,
	0x89, 0x4f, 0x0b, 0x7f, 0xc1, 0xa3, 0xcb, 0x5c, 0x7f, 0xa5, 0xc1, 0xee, 0x5c, 0xb4, 0x8d, 0x0a,
	0x37, 0x5e, 0x0e, 0x7e, 0xec, 0x5c, 0x1f, 0x1a, 0xd5, 0xf9, 0xf3, 0x91, 0xa1, 0xf1, 0x13, 0x5a,
	0xad, 0xb6, 0x38, 0x01, 0x66, 0x2f, 0xfc, 0x6c, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0xc2, 0x12, 0x5b, 0x37, 0x0b, 0x00, 0x00,
}