// const url = "http://localhost:8000";

export function getRam() {
    return fetch(`/insertRam`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}