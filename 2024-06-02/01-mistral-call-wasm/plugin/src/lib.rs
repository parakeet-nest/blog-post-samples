use extism_pdk::*;

#[plugin_fn]
pub fn hello(name: String) -> FnResult<String> {
    Ok(format!("👋 Hello, {}!", name))
}

#[plugin_fn]
pub fn hey(name: String) -> FnResult<String> {
    Ok(format!("🙂 Hey, {}! What's up?", name))
}

