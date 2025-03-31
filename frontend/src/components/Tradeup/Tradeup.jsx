import { useState, useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"
import { textMap } from "../../constants/rarity"
import CountdownTimer from "./CountdownTimer"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId
  const [tradeup, setTradeup] = useState({})
  const [loading, setLoading] = useState(true)
  const textColor = textMap[tradeup.rarity]

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:8080/v1/tradeups/" + tradeupId)

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
    <div className="flex flex-col items-center gap-2 mt-5">
      <div className="flex items-center gap-5">
        <span className={`font-bold text-2xl ${textColor}`}>{tradeup.rarity}</span>
        <span className="font-bold">—</span>
        <span className="font-bold text-2xl text-info">{tradeup.status}</span>
      </div>
      {tradeup.skins.length === 10 && tradeup.status === 'Active' && (
        <div className="font-bold text-lg">Tradeup Closes In: <CountdownTimer stopTime={tradeup.stopTime} /></div>
      )}
      <TradeupGrid tradeupId={tradeupId} rarity={tradeup.rarity} skins={tradeup.skins} status={tradeup.status} />
    </div>
  )
}

export default Tradeup
