import { useEffect, useState } from "react"
import TradeupRow from "../components/Tradeup/TradeupRow"
import { btnMap } from "../constants/rarity"

function ActiveTradeups() {
  const [tradeups, setTradeups] = useState([])
  const [selectedRarity, setSelectedRarity] = useState("All")
  const [rarityOptions, setRarityOptions] = useState(["All"])

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

      const rarityOrder = [
        "Consumer",
        "Industrial",
        "Mil-Spec",
        "Restricted",
        "Classified"
      ]

      // Extract unique status options
      const uniqueRarities = ["All", ...new Set(
        rarityOrder.filter(rarity =>
          newData.some(t => t.rarity === rarity)
        )
      )]
      setRarityOptions(uniqueRarities)
    }
    return () => {
      console.log("Closing SSE connection")
      eventSource.close()
    }
  }, [])

  const filteredTradeups = selectedRarity === "All"
    ? tradeups
    : tradeups.filter(t => t.rarity === selectedRarity)

  const handleReset = () => {
    setSelectedRarity("All")
  }

  return (
    <div className="flex flex-col items-center gap-2 mt-3">
      <h1 className="text-warning text-3xl font-bold">Active Trade Ups</h1>

      {/*Status Filter Dropdown */}
      <div className="filter mb-4 flex gap-1">
        <input
          className="btn btn-soft filter-reset"
          type="radio"
          name=""
          aria-label="×"
          onClick={handleReset}
        />
        {rarityOptions.filter(r => r !== "All").map(rarity => (
          <input
            key={rarity}
            className={`btn btn-soft ${btnMap[rarity] || ''}`}
            type="radio"
            name="rarity"
            aria-label={rarity}
            checked={selectedRarity === rarity}
            onChange={() => setSelectedRarity(rarity)}
          />
        ))}
      </div>

      {filteredTradeups && filteredTradeups.length > 0 ? (
        filteredTradeups.map(t => (
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
        <div>No active tradeups found for the selected status</div>
      )}
    </div>
  )
}

export default ActiveTradeups
