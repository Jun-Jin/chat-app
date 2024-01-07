import { useCallback, useEffect, useRef, useState } from 'react'
import './App.css'

function App() {
  const [message, setMessage] = useState<string[]>([])
  function appendMessage(message: string) {
    setMessage((messages) => [...messages, message])
  }
  const webSocketRef = useRef<WebSocket | null>(null)

  useEffect(() => {
    const webSocket = new WebSocket('ws://localhost:8080/ws')
    webSocketRef.current = webSocket

    webSocket.onopen = () => {
      console.log('WebSocket Client Connected')
    }
    webSocket.onmessage = (message) => {
      console.log('message', message)
      appendMessage(message.data)
    }
    webSocket.onclose = () => {
      console.log('WebSocket Client Disconnected')
    }
    return () => {
      webSocket.close()
    }
  }, [])

  const [input, setInput] = useState('')
  const submit = useCallback(() => {
    if (webSocketRef.current) {
      webSocketRef.current.send(input)
    }
  }, [input])

  return (
    <div className="App">
      <form onSubmit={(e) => e.preventDefault()}>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button onClick={submit}>Send</button>
      </form>
      <div>
        {message.map((m, i) => (
          <div key={i}>{m}</div>
        ))}
      </div>
    </div>
  )
}

export default App
