import { useState, useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId
  const [tradeup, setTradeup] = useState({})
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:8080/api/tradeups/" + tradeupId)

    eventSource.onopen = () => {
      console.log("Tradeup SSE connection opened successfully")
    }

    eventSource.onerror = (error) => {
      console.error("Tradeup SSE connection error:", error)
    }

    eventSource.onmessage = (event) => {
      const data = JSON.parse(event.data)
      setTradeup(data)
      setLoading(false)
    }
    return () => {
      console.log("Closing Tradeup SSE connection")
      eventSource.close()
    }
  }, [])

  if (loading) {
    return (
      <div className="flex flex-col items-center mt-5">
        <div className="loading-spinner"></div>
      </div>
    )
  }

  return (
    <div className="flex flex-col items-center mt-5">
      <div className="font-bold text-xl">Trade Up Status:&nbsp;
        <span className="text-info">{tradeup.status}</span>
      </div>
      <TradeupGrid tradeupId={tradeupId} rarity={tradeup.rarity} skins={tradeup.skins} />
    </div>
  )
}

export default Tradeup
