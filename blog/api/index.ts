import axios from 'axios';

const request = axios.create({
    baseURL: "http://127.0.0.1:8000/api/v1/",
    timeout: 3000
})

export async function get (url:string): Promise<any>{
    return request.get(url)
}


export async function post (url:string, data: any): Promise<any>{
    return request.post(url, data)
}

export function confirm (response:any) {
    const { data } = response
    if (data.state.code === 200) {
        return data.data
    } else {
        return false
    }
}
