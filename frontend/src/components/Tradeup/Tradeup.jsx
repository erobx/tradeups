import { useState, useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"
import { textMap } from "../../constants/rarity"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId
  const [tradeup, setTradeup] = useState({})
  const [loading, setLoading] = useState(true)
  const textColor = textMap[tradeup.rarity]

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
      <span className={`font-bold text-2xl ${textColor}`}>{tradeup.rarity}</span>
      <div className="font-bold text-xl">Status:&nbsp;
        <span className="text-info">{tradeup.status}</span>
      </div>
      {tradeup.skins.length === 10 && (
        <div>Tradeup Locks In:{}</div>
      )}
      <TradeupGrid tradeupId={tradeupId} rarity={tradeup.rarity} skins={tradeup.skins} />
    </div>
  )
}

export default Tradeup
