import { useEffect } from "react"
import { useParams } from "react-router"

function Tradeup() {
  const params = useParams()
  const tradeupId = params.tradeupId

  useEffect(() => {
    console.log(tradeupId)
  }, [])

  return (
    <div>
      {tradeupId}
    </div>
  )
}

export default Tradeup
