import useUser from "../../stores/userStore"
import RecentTradeups from "./RecentTradeups"
import Stats from "./Stats"

function Dashboard() {
  const { user, setUser, setBalance } = useUser()

  return (
    <div className="flex gap-42">
      <div className="flex-auto">
        <RecentTradeups user={user} />
      </div>
      <div className="divider divider-horizontal"></div>
      <div className="flex-auto">
        <Stats user={user} />
      </div>
    </div>
  )
}

export default Dashboard
