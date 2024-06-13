import {useEffect, useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {createProcess, deleteProcess, getRam} from "./api/Endpoint.jsx";

function App() {

    const [ percentRam, setPercentRam ] = useState(0);
    const [ actionButton, setActionButton ] = useState("Crear Proceso");
    const [ pid, setPid ] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            (
                async () => {
                    try {
                        const req = await getRam();
                        const res = await req.json();
                        console.log(res);
                        setPercentRam(res.percent);
                    } catch (e) {
                        console.log(e);
                    }
                }
            )();
        }, 500);

        return () => clearInterval(interval);

    }, [percentRam]);

    const createProc = async () => {
        try {
            const req = await createProcess();
            const res = await req.json();
            setPid(res.pid);
            setActionButton("Matar Proceso");
        } catch (e) {
            console.log(e);
        }
    }

    const deleteProc = async () => {
        try {
            const req = await deleteProcess(pid);
            const res = await req.json();
            console.log(res);
            setActionButton("Crear Proceso");
        } catch (e) {
            console.log(e);
        }
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
      <h1>RAM: {percentRam}%</h1>
      <div className="card">
          { actionButton === "Crear Proceso" ? (
                  <button onClick={createProc}>
                      { actionButton }
                  </button>
              )
              : (
                  <button onClick={deleteProc}>
                      {actionButton}
                  </button>
              )}
          <p>
              {pid}
          </p>
      </div>
        <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
