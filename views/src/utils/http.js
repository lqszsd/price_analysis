import axios from 'axios';
//let BaseURl = 'http://127.0.0.1:9090';   
let BaseURl = 'http://127.0.0.1:9090'
// get post请求封装
export function get(url, param) {
    return new Promise((resolve, reject) => {
        console.log(BaseURl + url);
        axios.get(BaseURl + url, {params: param}).then(response => {
            resolve(response.data)
        }, err => {
            reject(err)
        }).catch((error) => {
            reject(error)
        })
    })
}               // 你的域名

export function post(url, params) {
    return new Promise((resolve, reject) => {
        axios.post(BaseURl + url, params).then(response => {
            resolve(response.data);
        }, err => {
            reject(err);
        }).catch((error) => {
            reject(error)
        })
    })
}
