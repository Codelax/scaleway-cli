// Tests have to run with go installed
// It will load go misc files in
// /usr/local/go/misc/wasm

import {describe, it, expect, afterAll, beforeAll} from 'vitest'

import '../wasm_exec_node.cjs'
import '../wasm_exec.cjs'
import {CLI, RunConfig} from '../cli'
import * as fs from 'fs'
import {Go} from "../wasm_exec";
import {loadWasmBinary} from "./utils";

const CLI_PACKAGE = 'scw'
const CLI_CALLBACK = 'cliLoaded'

describe('With wasm CLI', async () => {
    let cli: CLI

    beforeAll(async () => {
        // @ts-ignore
        cli = await loadWasmBinary('./cli.wasm') as CLI
    })

    const run = async (expected: string | RegExp, command: string[], runCfg: RunConfig | null = null) => {
        if (runCfg === null) {
            runCfg = {
                jwt: "",
            }
        }

        const resp = await cli.run(runCfg, command)
        expect(resp.exitCode).toBe(0)
        expect(resp.stdout).toMatch(expected)
    }

    const runWithError = async (expected: string | RegExp, command: string[], runCfg: RunConfig | null = null) => {
        if (runCfg === null) {
            runCfg = {
                jwt: "",
            }
        }
        const resp = await cli.run(runCfg, command)
        expect(resp.exitCode).toBeGreaterThan(0)
        expect(resp.stderr).toMatch(expected)
    }

    const complete = async (expected: string[], command: string[], runCfg: RunConfig | null = null) => {
        if (runCfg === null) {
            runCfg = {
                jwt: "",
            }
        }
        let toComplete = command.pop() || ""

        const suggestions = await cli.complete({
            jwt: runCfg.jwt,
            leftWords: command,
            rightWords: [],
            selectedWord: toComplete
        })
        expected.forEach(suggestion => expect(suggestions).toContain(suggestion))
    }

    it('can run cli commands', async () => run(/profile.*default/, ['info']))

    it('can run help', async () => runWithError(/USAGE:\n.*scw <command>.*/, []))

    it('can use jwt', async () => runWithError(/.*denied authentication.*invalid JWT.*/, ['instance', 'server', 'list']))

    it('can complete', async () => complete(['server', 'image', 'volume'], ['instance', '']))

    afterAll(async () => {
        try {
            await cli.stop()
            go._resume()
        } catch (e) {
            console.log(e)
        }
    })
})
