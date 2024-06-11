import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {getRam} from "./api/Endpoint.jsx";

function App() {

    const [ percentRam, setPercentRam ] = useState("");

    const insertPercentRam = async () => {
        const req = await getRam();
        const res = await req.json();
        console.log(res);
        setPercentRam(res);
    }

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>RAM: {percentRam}</h1>
      <div className="card">
        <button onClick={insertPercentRam}>
            Presiona para ver RAM
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
