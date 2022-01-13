import {onDestroy} from "svelte";

export function zfill(number: number, size: number): string {
    let num: string = number.toString()
    while (num.length < size) num = "0" + num
    return num
}

export function getTime(d: Date): string {
    return zfill(d.getHours(), 2) + ':' + zfill(d.getMinutes(), 2)
}

export function getDate(d: Date): string {
    return zfill(d.getFullYear(), 2) + '-' + zfill(d.getMonth() + 1, 2) + '-' + zfill(d.getDate(), 2)
}

export function getNextDay(d: Date): Date {
    return new Date(
        d.getFullYear(), d.getMonth(), d.getDate() + 1,
        0, 0, 0, 0
    )
}

export function getLastDay(d: Date): Date {
    return new Date(
        d.getFullYear(), d.getMonth(), d.getDate() - 1,
        0, 0, 0, 0
    )
}

export function genHours(d: Date): Array<Date> {
    const nextDate: Date = getNextDay(d)

    // Used in loop
    let tmpDate: Date = new Date(
        d.getFullYear(), d.getMonth(), d.getDate(),
        0, 0, 0, 0
    )

    let hours: Array<Date> = []

    while (tmpDate.getDate() !== nextDate.getDate()) {
        hours.push(new Date(
            tmpDate.getFullYear(), tmpDate.getMonth(), tmpDate.getDate(),
            tmpDate.getHours(), tmpDate.getMinutes(), 0
        ))

        //tmpDate.setHours(tmpDate.getHours() + 1)
        tmpDate.setMinutes(tmpDate.getMinutes() + 60)
    }

    return hours
}


export function handlePageSubmit(endpointUrl: string, method: string, values: object, successFunc: Function = function (o: object) {
}) {
    fetch(endpointUrl, {
        method: method,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(values)
    }).then((res) => {
        // Handle internal server error
        if (res.status === 500) {
            alert("internal server error :(")
            return
        }

        if (res.status === 404) {
            alert("not found :(")
            return
        }

        return res

    }).then((res) => {
        if (res === undefined) {
            return
        }

        res.json().then(j => {
            if (res.status === 400) {
                // Invalid values in form field(s)
                console.log("errored")

                j.forEach(item => {
                    console.log(item["field"])
                    console.log(item["msg"])
                })

                return
            }

            // Handle success
            successFunc(j)
        })

    }).catch(ex => {
        alert(ex)
    })

}

export function getWeekNumber(d: Date): string {
    d = new Date(Date.UTC(d.getFullYear(), d.getMonth(), d.getDate()));
    d.setUTCDate(d.getUTCDate() + 4 - (d.getUTCDay() || 7));
    const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1));
    const weekNo: number = Math.ceil(
        ((d.getTime() - yearStart.getTime()) / 86400000 + 1) / 7
    );

    return zfill(weekNo, 2)
}


export function isNumber(value: string | number): boolean {
    return ((value != null) &&
        (value !== '') &&
        !isNaN(Number(value.toString())));
}

export interface Level {
    Name: string
    Show: boolean
    Worst: string
}

export function onInterval(callback, milliseconds) {
    const interval = setInterval(callback, milliseconds);

    onDestroy(() => {
        clearInterval(interval);
    });
}