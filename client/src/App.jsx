function App() {
  return (
    <div>
      <h1>Hello go</h1>
      <button
        onClick={async () => {
          const response = await fetch("/users");
          const data = await response.json();
          console.log(data);
        }}
      >
        Obtener datos
      </button>
    </div>
  );
}

export default App;
