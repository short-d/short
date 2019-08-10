interface Extensions {
    code: string
}

export interface GraphQlError {
    extensions: Extensions
    message: string
    path: string[]
}