export interface questioni18n {
    Title: string
    Description: string
}

export interface answeri18n {
    Description: string
}

export interface question {
    Question: questioni18n
    Answer: Array<answeri18n>
}