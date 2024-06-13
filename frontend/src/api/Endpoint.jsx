// const url = "http://localhost:8000";

export function getRam() {
    return fetch(`/getRam`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}

export function createProcess() {
    return fetch(`/insertProcess`, {
        method: 'GET',
        headers: {'Content-Type': 'application/json'},
    })
}

export function deleteProcess(pid) {
    return fetch(`/delProcess?pid=${pid}`, {
        method: 'GET',
        headers: {'Content-Type': 'application/json'},
    })
}