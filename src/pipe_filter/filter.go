package pipefilter

//Request输入filter
type Request interface{}

//Response输出filter
type Response interface{}

//Pipe-Filter structure
type Filter interface {
	Process(data Request)(Response, error)
}
