
function Balance({ balance }) {
  return (
    <button className="btn badge-lg btn-success mr-1.5">
      ${parseFloat(balance).toFixed(2)}
    </button>
  )
}

export default Balance
