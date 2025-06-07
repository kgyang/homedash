import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Env from './env.jsx'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>Home Dashboard</h1>
      <Env />
      <p className="read-the-docs">
        Home dash board
      </p>
    </>
  )
}

export default App
