package node

import (
    "sync"
)

type Plugin struct {
    Node   *Node
    Name   string
    Events pluginEvents
    wg     *sync.WaitGroup
}

func NewPlugin(name string, callback Callback, callbacks ...Callback) *Plugin {
    plugin := &Plugin{
        Name: name,
        Events: pluginEvents{
            Configure: &callbackEvent{make(map[uintptr]Callback)},
            Run:       &callbackEvent{make(map[uintptr]Callback)},
        },
    }

    if len(callbacks) >= 1 {
        plugin.Events.Configure.Attach(callback)
        for _, callback = range callbacks[:len(callbacks)-1] {
            plugin.Events.Configure.Attach(callback)
        }

        plugin.Events.Run.Attach(callbacks[len(callbacks)-1])
    } else {
        plugin.Events.Run.Attach(callback)
    }

    return plugin
}

func (plugin *Plugin) LogSuccess(message string) {
    plugin.Node.LogSuccess(plugin.Name, message)
}

func (plugin *Plugin) LogInfo(message string) {
    plugin.Node.LogInfo(plugin.Name, message)
}

func (plugin *Plugin) LogWarning(message string) {
    plugin.Node.LogWarning(plugin.Name, message)
}

func (plugin *Plugin) LogFailure(message string) {
    plugin.Node.LogFailure(plugin.Name, message)
}