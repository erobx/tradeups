import { useEffect } from "react"
import { useParams } from "react-router"
import TradeupGrid from "./TradeupGrid"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId

  useEffect(() => {
    console.log(tradeupId)
  }, [])

  return (
    <div className="flex flex-col items-center">
      <TradeupGrid />
    </div>
  )
}

export default Tradeup
