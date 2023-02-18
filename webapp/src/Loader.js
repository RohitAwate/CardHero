const Loader = () => {
    const style = {
        display: "flex", alignItems: "center", justifyContent: "center",
    };

    return <img style={style}
                className="loader"
                src="/loaders/ellipsis-green.svg"
                alt="Loading..."
    />;
};

export default Loader;