import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [apiStatus, setApiStatus] = useState('Checking...')

  useEffect(() => {
    // Test API connection
    fetch('/health')
      .then(response => response.json())
      .then(data => setApiStatus(data.status))
      .catch(() => setApiStatus('API not available'))
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        <h1>SEC Baseball Hub</h1>
        <p className="coming-soon">Coming Soon</p>
        <div className="api-status">
          <p>API Status: {apiStatus}</p>
        </div>
      </header>
    </div>
  )
}

export default App
