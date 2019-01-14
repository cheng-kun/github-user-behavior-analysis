HTTP GET

HOST:52.14.21.138:2222
ENDPOINT:/api/?days={days}&language={language}

RESPONSE:JSON list

EXAMPLE

GET:52.14.21.138:2222/api/?days=3&language=Python

RESPONSE:

[4021.0,4009.0,3536.0]

days:预测天数，预测天数为n，返回的list的lenth为n
language:['C#', 'C++', 'CSS', 'HTML', 'Java', 'JavaScript', 'PHP', 'Python', 'Ruby', 'TypeScript', 'Perl', 'C']
