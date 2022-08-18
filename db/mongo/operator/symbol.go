/*
Create: 2022/8/19
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package operator
package operator

// mongo的nosql操作符号

// Comparison
const (
	Eq  = "$eq"
	Gt  = "$gt"
	Gte = "$gte"
	In  = "$in"
	Lt  = "$lt"
	Lte = "$lte"
	Ne  = "$ne"
	Nin = "$nin"
)

// Logical
const (
	And = "$and"
	Not = "$not"
	Nor = "$nor"
	Or  = "$or"
)

// Element
const (
	Exists = "$exists"
	Type   = "$type"
)

// Evaluation
const (
	Expr       = "$expr"
	JSONSchema = "$jsonSchema"
	Mod        = "$mod"
	Regex      = "$regex"
	Text       = "$text"
	Where      = "$where"
)

// Geo spatial
const (
	GeoIntersects = "$geoIntersects"
	GeoWithin     = "$geoWithin"
	Near          = "$near"
	NearSphere    = "$nearSphere"
)

// Array
const (
	All       = "$all"
	ElemMatch = "$elemMatch"
	Size      = "$size"
)

// Bitwise
const (
	BitsAllClear = "$bitsAllClear"
	BitsAllSet   = "$bitsAllSet"
	BitsAnyClear = "$bitsAnyClear"
	BitsAnySet   = "$bitsAnySet"
)

// Comments
const (
	Comment = "$comment"
)

// Projection operators
const (
	Dollar = "$"
	// ElemMatch = "$elemMatch" // Declared
	Meta  = "$meta"
	Slice = "$slice"
)

// Modifiers
const (
	// Comment = "$comment" // Declared
	Explain = "$explain"
	Hint    = "$hint"
	// Max     = "$max" // Declared
	MaxTimeMS = "$maxTimeMS"
	// Min       = "$min" // Declared
	OrderBy     = "$orderby"
	Query       = "$query"
	ReturnKey   = "$returnKey"
	ShowDiskLoc = "$showDiskLoc"
)

// Sort Order
const (
	Natural = "$natural"
)

// Fields
const (
	CurrentDate = "$currentDate"
	Inc         = "$inc"
	Min         = "$min"
	Max         = "$max"
	Mul         = "$mul"
	Rename      = "$rename"
	Set         = "$set"
	SetOnInsert = "$setOnInsert"
	Unset       = "$unset"
)

// Array Operators
const (
	// $: Act as a modifier
	// $[]: Act as a modifier
	// $[<identifier>]: Act as a modifier

	AddToSet = "$addToSet"
	Pop      = "$pop"
	Pull     = "$pull"
	Push     = "$push"
	PullAll  = "$pullAll"
)

// Array modifiers
const (
	Each     = "$each"
	Position = "$position"
	// Slice    = "$slice" // Declared
	Sort = "$sort"
)

// Array bitwise
const (
	Bit = "$bit"
)

// Collection aggregation Stages
const (
	AddFields      = "$addFields"
	Bucket         = "$bucket"
	BucketAuto     = "$bucketAuto"
	CollStats      = "$collStats"
	Count          = "$count"
	Facet          = "$facet"
	GeoNear        = "$geoNear"
	GraphLookup    = "$graphLookup"
	Group          = "$group"
	IndexStats     = "$indexStats"
	Limit          = "$limit"
	ListSessions   = "$listSessions"
	Lookup         = "$lookup"
	Match          = "$match"
	Merge          = "$merge"
	Out            = "$out"
	PlanCacheStats = "$planCacheStats"
	Project        = "$project"
	Redact         = "$redact"
	ReplaceRoot    = "$replaceRoot"
	ReplaceWith    = "$replaceWith"
	Sample         = "$sample"
	// Set            = "$set" // Declared
	Skip = "$skip"
	// Sort           = "$sort" // Declared
	SortByCount = "$sortByCount"
	// Unset          = "$unset" // Declared
	Unwind = "$unwind"
)

// DB Aggregate stages
const (
	CurrentOp         = "$currentOp"
	ListLocalSessions = "$listLocalSessions"
)

// Arithmetic Expression Operators
const (
	Abs    = "$abs"
	Add    = "$add"
	Ceil   = "$ceil"
	Divide = "$divide"
	Exp    = "$exp"
	Floor  = "$floor"
	Ln     = "$ln"
	Log    = "$log"
	Log10  = "$log10"
	// Mod      = "$mod" // Declared
	Multiply = "$multiply"
	Pow      = "$pow"
	Round    = "$round"
	Sqrt     = "$sqrt"
	Subtract = "$subtract"
	Trunc    = "$trunc"
)

// Array Expression Operators
const (
	ArrayToObject = "$arrayToObject"
	ConcatArrays  = "$concatArrays"
	Filter        = "$filter"
	// In            = "$in" // Declared
	IndexOfArray  = "$indexOfArray"
	IsArray       = "$isArray"
	Map           = "$map"
	ObjectToArray = "$objectToArray"
	Range         = "$range"
	Reduce        = "$reduce"
	ReverseArray  = "$reverseArray"
	// Size          = "$size" // Declared
	// Slice         = "$slice" // Declared
	Zip = "$zip"
)

// Boolean Expression Operators
const (
// And = "$and" // Declared
// Not = "$not" // Declared
// Or  = "$or" // Declared
)

// Comparison Expression Operators
const (
	Cmp = "$cmp"
	//Eq  = "$eq" // Declared
	//Gt  = "$gt" // Declared
	//Gte = "$gte" // Declared
	//Lt  = "$lt" // Declared
	//Lte = "$lte" // Declared
	//Ne  = "$ne" // Declared
)

// Conditional Expression Operators
const (
	Cond   = "$cond"
	IfNull = "$ifNull"
	Switch = "$switch"
)

// Date Expression Operators
const (
	DateFromParts  = "$dateFromParts"
	DateFromString = "$dateFromString"
	DateToParts    = "$dateToParts"
	DateToString   = "$dateToString"
	DayOfMonth     = "$dayOfMonth"
	DayOfWeek      = "$dayOfWeek"
	DayOfYear      = "$dayOfYear"
	Hour           = "$hour"
	IsoDayOfWeek   = "$isoDayOfWeek"
	IsoWeek        = "$isoWeek"
	IsoWeekYear    = "$isoWeekYear"
	Millisecond    = "$millisecond"
	Minute         = "$minute"
	Month          = "$month"
	Second         = "$second"
	ToDate         = "$toDate"
	Week           = "$week"
	Year           = "$year"
)

// Literal Expression Operator
const (
	Literal = "$literal"
)

// Object Expression Operators
const (
	MergeObjects = "$mergeObjects"
	// ObjectToArray = "$objectToArray" // Declared
)

// Set Expression Operators
const (
	AllElementsTrue = "$allElementsTrue"
	AnyElementTrue  = "$anyElementTrue"
	SetDifference   = "$setDifference"
	SetEquals       = "$setEquals"
	SetIntersection = "$setIntersection"
	SetIsSubset     = "$setIsSubset"
	SetUnion        = "$setUnion"
)

// String Expression Operators
const (
	Concat = "$concat"
	// DateFromString = "$dateFromString" // Declared
	// DateToString   = "$dateToString" // Declared
	IndexOfBytes = "$indexOfBytes"
	IndexOfCP    = "$indexOfCP"
	Ltrim        = "$ltrim"
	RegexFind    = "$regexFind"
	RegexFindAll = "$regexFindAll"
	RegexMatch   = "$regexMatch"
	Rtrim        = "$rtrim"
	Split        = "$split"
	StrLenBytes  = "$strLenBytes"
	StrLenCP     = "$strLenCP"
	Strcasecmp   = "$strcasecmp"
	Substr       = "$substr"
	SubstrBytes  = "$substrBytes"
	SubstrCP     = "$substrCP"
	ToLower      = "$toLower"
	ToString     = "$toString"
	Trim         = "$trim"
	ToUpper      = "$toUpper"
)

// Text Expression Operator
const (
// Meta = "$meta" // Declared
)

// Trigonometry Expression Operators
const (
	Sin              = "$sin"
	Cos              = "$cos"
	Tan              = "$tan"
	Asin             = "$asin"
	Acos             = "$acos"
	Atan             = "$atan"
	Atan2            = "$atan2"
	Asinh            = "$asinh"
	Acosh            = "$acosh"
	Atanh            = "$atanh"
	DegreesToRadians = "$degreesToRadians"
	RadiansToDegrees = "$radiansToDegrees"
)

// Type Expression Operators
const (
	Convert = "$convert"
	ToBool  = "$toBool"
	//ToDate     = "$toDate" // Declared
	ToDecimal  = "$toDecimal"
	ToDouble   = "$toDouble"
	ToInt      = "$toInt"
	ToLong     = "$toLong"
	ToObjectID = "$toObjectId"
	//ToString   = "$toString" // Declared
	//Type       = "$type" // Declared
)

// Accumulators ($group)
const (
	// AddToSet     = "$addToSet" // Declared
	Avg   = "$avg"
	First = "$first"
	Last  = "$last"
	// Max          = "$max" // Declared
	// MergeObjects = "$mergeObjects" // Declared
	// Min          = "$min" // Declared
	// Push         = "$push" // Declared
	StdDevPop  = "$stdDevPop"
	StdDevSamp = "$stdDevSamp"
	Sum        = "$sum"
)

// Accumulators (in Other Stages)
const (
// Avg        = "$avg" // Declared
// Max        = "$max" // Declared
// Min        = "$min" // Declared
// StdDevPop  = "$stdDevPop" // Declared
// StdDevSamp = "$stdDevSamp" // Declared
// Sum        = "$sum" // Declared
)

// Variable Expression Operators
const (
	Let = "$let"
)
