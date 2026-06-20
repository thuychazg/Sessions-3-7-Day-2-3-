import {useEffect,useState} from "react"
import {
    getAssets,
    createAsset,
    deleteAsset,
    scanAsset
}
from "./api"

import "./index.css"



function App(){


const [assets,setAssets]=useState([])

const [name,setName]=useState("")

const [type,setType]=useState("domain")

const [result,setResult]=useState(null)



async function loadAssets(){

    const data =
    await getAssets()


    setAssets(
        data.data || []
    )

}



useEffect(()=>{

    loadAssets()

},[])




async function add(){


    if(!name)
        return


    await createAsset({

        name:name,
        type:type

    })


    setName("")

    loadAssets()

}





async function remove(id){

    await deleteAsset(id)

    loadAssets()

}




async function scan(){


    const data =
    await scanAsset()


    setResult(data)

}





return (

<div className="container">


<h1>
Asset Scanner Dashboard
</h1>


<div className="card">


<h2>
Create Asset
</h2>


<input

placeholder="google.com"

value={name}

onChange={
e=>setName(e.target.value)
}

/>


<select

value={type}

onChange={
e=>setType(e.target.value)
}

>

<option value="domain">
domain
</option>


<option value="ip">
ip
</option>


<option value="service">
service
</option>


</select>


<button onClick={add}>
Add Asset
</button>


</div>





<div className="card">


<h2>
Statistics
</h2>


<p>
Total assets:
<b>
 {assets.length}
</b>
</p>


</div>





<div className="card">


<h2>
Assets
</h2>



{

assets.map(a=>

<div className="asset"
key={a.id}>


<span>

{a.name}

&nbsp;

({a.type})

</span>



<button
onClick={
()=>remove(a.id)
}
>

Delete

</button>


</div>


)

}


</div>






<div className="card">


<h2>
Scanner
</h2>


<button onClick={scan}>

Start Scan

</button>



{

result &&

<pre>

{
JSON.stringify(
result,
null,
2
)
}

</pre>

}


</div>





</div>

)


}


export default App
