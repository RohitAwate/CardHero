const Loader = (props) => {
    const width = props.size ? props.size : 100;

    const style = {
        display: "flex", alignItems: "center", justifyContent: "center", width: width
    };

    return <img style={style}
                className="loader"
                src="/loaders/ellipsis-green.svg"
                alt="Loading..."
    />;
};

export default Loader;