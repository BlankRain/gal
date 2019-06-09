# gal
Graph Algorithm  Language

## Samples

```
package   wo.le.ge.ca.ca.ca

NodeType Page{
    URL: string @index(exact, fulltext) @count  @required @updated   @filter(StartWith('hello')),
    Name:   string @required(false)    ,
}@Query{
    has(website){
        uid
        url
        name
    }  
}

EdgeType  referTo(s:Person,t:Person) @filter(has(s.Name)) @reverse {
    id,
    lable: s.Name,
}


from Graph([Page],[referTo])  g
make function  PageRank()  {
    print(g)
}

from Graph([Page],[referTo])  g  
make function  ShortestPath(sourePage Page ,endPage Page) Path([Page]){
    print(g)
}


```

## Support Feature
- package 
- NodeType
- EdgeType
- from graph ... make function
- /****/ as comment