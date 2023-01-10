export interface ApiRes<T> {
    code: number,
    msg?: string,
    data: T
}
