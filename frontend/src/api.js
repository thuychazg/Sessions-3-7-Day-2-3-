const API = import.meta.env.VITE_API_URL


export async function getAssets(){

    const res =
    await fetch(
        `${API}/assets`
    )

    return res.json()
}



export async function createAsset(asset){

    const res =
    await fetch(
        `${API}/assets`,
        {
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify(asset)
        }
    )

    return res.json()
}



export async function deleteAsset(id){

    return fetch(
        `${API}/assets?ids=${id}`,
        {
            method:"DELETE"
        }
    )

}



export async function scanAsset(){

    const res =
    await fetch(
        `${API}/scan`,
        {
            method:"POST",
            headers:{
                "Content-Type":"application/json"
            },
            body:JSON.stringify({
                scan_type:"tech"
            })
        }
    )

    return res.json()
}
