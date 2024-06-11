---
outline: false
---

# 🤔 FAQs

[[toc]]

## ❓ Questions

### How do I upload already existing files?

You can simply drag & drop existing files from your local browser into the editor tree-view area to upload them.

### Ho do I interact with Arkanum from the integrated terminal?

You can use the the `code-server` command from the integrated terminal sessions like bash or pwsh.
To reuse the existing session / window and open the current location just run `code-server -r .`

> [!WARNING] ⚠️ WARNING
> The command `code-server` is only available as user, but not as root.

### How should I start developing? What are my first steps?

We recommend you take a look at our maintained templates:

- [Baseline](https://gitea.ocram85.com/Templates/Baseline)

### How do I install external resources?

To install external resources you can pass the already set proxy configuration into root a session. Therefore use
the `sudo -E` command. Don't forget the `-E` parameter to preserve the proxy config.

> [!WARNING] ⚠️ WARNING
> Additionally we need to restrict the internet access. This causes us run a proxy with an active whitelist filter.
> So you could end up getting an `http 403` error.

### Whats the root password for my Arkanum instance?

The root password is provided by the container environment variables. Just search for the value of `SUDO_PASSWORD`.
