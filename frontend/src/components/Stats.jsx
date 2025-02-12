

function Stats() {
  return (
    <div className="stats bg-base-300 shadow-md">
      <div className="stat">
        <div className="stat-title">Account balance</div>
        <div className="stat-value">$420.69</div>
        <div className="stat-actions">
          <button className="btn btn-xs btn-success">Add funds</button>
        </div>
      </div>

      <div className="stat">
        <div className="stat-title">Recent winnings</div>
        <div className="stat-value">$200</div>
        <div className="stat-actions">
          <button className="btn btn-xs btn-error">Withdrawal</button>
        </div>
      </div>
    </div>
  )
}

export default Stats
