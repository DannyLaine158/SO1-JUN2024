const url = "http://localhost:8000";

export function getRam() {
    return fetch(`${url}/insertRam`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}