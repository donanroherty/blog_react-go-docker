import logo from "./logo.svg"
import "./App.css"

import { useState } from "react"

function App() {
  const [time, setTime] = useState(null)

  const handleGetTime = async () => {
    const res = await fetch("/api/time")
    const json = await res.json()
    setTime(json.server_time)
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <button onClick={handleGetTime}>What time is it?</button>
        {time && <p>The time is: {time}</p>}
      </header>
    </div>
  )
}

export default App
