module github.com/conflabermits/trusty_web_server

go 1.18

replace github.com/conflabermits/trusty_web_server/httpfunctions => ./pkg

require github.com/conflabermits/trusty_web_server/httpfunctions v0.1.0-alpha
