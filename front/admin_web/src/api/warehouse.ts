import { httpHeader, AxiosPromise } from "./api"

export let saveWarehouse = (name: string): AxiosPromise => {
    let params = {
        name: name
    }
    return httpHeader("POST", "/warehouse/save", params, {})
}

export let modifyWarehouse = (id: number, name: string): AxiosPromise => {
    let params = {
        id: id,
        name: name
    }
    return httpHeader("POST", "/warehouse/modify", params, {})
}

export let getWarehouseList = (): AxiosPromise => {
    return httpHeader("GET", "/warehouse/get-list", {}, {})
}