import { httpHeader, AxiosPromise } from "./api"

export let savePlatfrom = (name: string): AxiosPromise => {
    let params = {
        name: name
    }
    return httpHeader("POST", "/platform/save", params, {})
}

export let modifyPlatfrom = (id: number, name: string): AxiosPromise => {
    let params = {
        id: id,
        name: name
    }
    return httpHeader("POST", "/platform/modify", params, {})
}

export let getPlatfromList = (): AxiosPromise => {
    return httpHeader("GET", "/platform/get-list", {}, {})
}