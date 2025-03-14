import { useEffect, useState } from "react"
import TradeupRow from "../components/Tradeup/TradeupRow"

function ActiveTradeups() {
  const [tradeups, setTradeups] = useState([])

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:8080/api/tradeups")

    eventSource.onopen = () => {
      console.log("SSE connection opened successfully")
    }

    eventSource.onerror = (error) => {
      console.error("SSE connection error:", error)
    }

    eventSource.onmessage = (event) => {
      const newData = JSON.parse(event.data)
      setTradeups(newData)
    }
    return () => {
      console.log("Closing SSE connection")
      eventSource.close()
    }
  }, [])

  return (
    <div className="flex flex-col items-center gap-2 mt-3">
      <h1 className="text-warning text-3xl font-bold">Active Trade Ups</h1>
      <div>Filter placeholder</div>
      {tradeups && tradeups.length > 0 ? (
        tradeups.map(t => (
          <TradeupRow
            key={t.id}
            id={t.id}
            players={t.players}
            rarity={t.rarity}
            skins={t.skins}
            status={t.status}
          />
        ))
      ) : (
        <div>No active tradeups found</div>
      )}
    </div>
  )
}

export default ActiveTradeups
