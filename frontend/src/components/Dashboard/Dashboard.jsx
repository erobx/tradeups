import useUser from "../../stores/userStore"
import RecentTradeups from "./RecentTradeups"
import Stats from "./Stats"

function Dashboard() {
  const { user, setUser, setBalance } = useUser()

  return (
    <div className="flex mb-5">
      <RecentTradeups user={user} />
      <div className="divider lg:divider-horizontal"></div>
      <Stats user={user} />
    </div>
  )
}

export default Dashboard
