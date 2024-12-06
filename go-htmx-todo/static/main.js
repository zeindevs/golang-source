function getDefaultExportFromCjs(ze) {
  return ze && ze.__esModule && Object.prototype.hasOwnProperty.call(ze, "default") ? ze.default : ze;
}
var htmx_min$1 = { exports: {} }, htmx_min = htmx_min$1.exports, hasRequiredHtmx_min;
function requireHtmx_min() {
  return hasRequiredHtmx_min || (hasRequiredHtmx_min = 1, function(module) {
    (function(ze, Qr) {
      module.exports ? module.exports = Qr() : ze.htmx = ze.htmx || Qr();
    })(typeof self < "u" ? self : htmx_min, function() {
      return function() {
        var Q = { onLoad: F, process: zt, on: de, off: ge, trigger: ce, ajax: Nr, find: C, findAll: f, closest: v, values: function(ze, Qr) {
          var Wr = dr(ze, Qr || "post");
          return Wr.values;
        }, remove: _, addClass: z, removeClass: n, toggleClass: $, takeClass: W, defineExtension: Ur, removeExtension: Br, logAll: V, logNone: j, logger: null, config: { historyEnabled: !0, historyCacheSize: 10, refreshOnHistoryMiss: !1, defaultSwapStyle: "innerHTML", defaultSwapDelay: 0, defaultSettleDelay: 20, includeIndicatorStyles: !0, indicatorClass: "htmx-indicator", requestClass: "htmx-request", addedClass: "htmx-added", settlingClass: "htmx-settling", swappingClass: "htmx-swapping", allowEval: !0, allowScriptTags: !0, inlineScriptNonce: "", attributesToSettle: ["class", "style", "width", "height"], withCredentials: !1, timeout: 0, wsReconnectDelay: "full-jitter", wsBinaryType: "blob", disableSelector: "[hx-disable], [data-hx-disable]", useTemplateFragments: !1, scrollBehavior: "smooth", defaultFocusScroll: !1, getCacheBusterParam: !1, globalViewTransitions: !1, methodsThatUseUrlParams: ["get"], selfRequestsOnly: !1, ignoreTitle: !1, scrollIntoViewOnBoost: !0, triggerSpecsCache: null }, parseInterval: d, _: t, createEventSource: function(ze) {
          return new EventSource(ze, { withCredentials: !0 });
        }, createWebSocket: function(ze) {
          var Qr = new WebSocket(ze, []);
          return Qr.binaryType = Q.config.wsBinaryType, Qr;
        }, version: "1.9.12" }, r = { addTriggerHandler: Lt, bodyContains: se, canAccessLocalStorage: U, findThisElement: xe, filterValues: yr, hasAttribute: o, getAttributeValue: te, getClosestAttributeValue: ne, getClosestMatch: c, getExpressionVars: Hr, getHeaders: xr, getInputValues: dr, getInternalData: ae, getSwapSpecification: wr, getTriggerSpecs: it, getTarget: ye, makeFragment: l, mergeObjects: le, makeSettleInfo: T, oobSwap: Ee, querySelectorExt: ue, selectAndSwap: je, settleImmediately: nr, shouldCancel: ut, triggerEvent: ce, triggerErrorEvent: fe, withExtensions: R }, w = ["get", "post", "put", "delete", "patch"], i = w.map(function(ze) {
          return "[hx-" + ze + "], [data-hx-" + ze + "]";
        }).join(", "), S = e("head"), q = e("title"), H = e("svg", !0);
        function e(ze, Qr) {
          return new RegExp("<" + ze + "(\\s[^>]*>|>)([\\s\\S]*?)<\\/" + ze + ">", Qr ? "gim" : "im");
        }
        function d(ze) {
          if (ze == null)
            return;
          let Qr = NaN;
          return ze.slice(-2) == "ms" ? Qr = parseFloat(ze.slice(0, -2)) : ze.slice(-1) == "s" ? Qr = parseFloat(ze.slice(0, -1)) * 1e3 : ze.slice(-1) == "m" ? Qr = parseFloat(ze.slice(0, -1)) * 1e3 * 60 : Qr = parseFloat(ze), isNaN(Qr) ? void 0 : Qr;
        }
        function ee(ze, Qr) {
          return ze.getAttribute && ze.getAttribute(Qr);
        }
        function o(ze, Qr) {
          return ze.hasAttribute && (ze.hasAttribute(Qr) || ze.hasAttribute("data-" + Qr));
        }
        function te(ze, Qr) {
          return ee(ze, Qr) || ee(ze, "data-" + Qr);
        }
        function u(ze) {
          return ze.parentElement;
        }
        function re() {
          return document;
        }
        function c(ze, Qr) {
          for (; ze && !Qr(ze); )
            ze = u(ze);
          return ze || null;
        }
        function L(ze, Qr, Wr) {
          var Gr = te(Qr, Wr), Yr = te(Qr, "hx-disinherit");
          return ze !== Qr && Yr && (Yr === "*" || Yr.split(" ").indexOf(Wr) >= 0) ? "unset" : Gr;
        }
        function ne(ze, Qr) {
          var Wr = null;
          if (c(ze, function(Gr) {
            return Wr = L(ze, Gr, Qr);
          }), Wr !== "unset")
            return Wr;
        }
        function h(ze, Qr) {
          var Wr = ze.matches || ze.matchesSelector || ze.msMatchesSelector || ze.mozMatchesSelector || ze.webkitMatchesSelector || ze.oMatchesSelector;
          return Wr && Wr.call(ze, Qr);
        }
        function A(ze) {
          var Qr = /<([a-z][^\/\0>\x20\t\r\n\f]*)/i, Wr = Qr.exec(ze);
          return Wr ? Wr[1].toLowerCase() : "";
        }
        function s(ze, Qr) {
          for (var Wr = new DOMParser(), Gr = Wr.parseFromString(ze, "text/html"), Yr = Gr.body; Qr > 0; )
            Qr--, Yr = Yr.firstChild;
          return Yr == null && (Yr = re().createDocumentFragment()), Yr;
        }
        function N(ze) {
          return /<body/.test(ze);
        }
        function l(ze) {
          var Qr = !N(ze), Wr = A(ze), Gr = ze;
          if (Wr === "head" && (Gr = Gr.replace(S, "")), Q.config.useTemplateFragments && Qr) {
            var Yr = s("<body><template>" + Gr + "</template></body>", 0), Jr = Yr.querySelector("template").content;
            return Q.config.allowScriptTags ? oe(Jr.querySelectorAll("script"), function(Zr) {
              Q.config.inlineScriptNonce && (Zr.nonce = Q.config.inlineScriptNonce), Zr.htmxExecuted = navigator.userAgent.indexOf("Firefox") === -1;
            }) : oe(Jr.querySelectorAll("script"), function(Zr) {
              _(Zr);
            }), Jr;
          }
          switch (Wr) {
            case "thead":
            case "tbody":
            case "tfoot":
            case "colgroup":
            case "caption":
              return s("<table>" + Gr + "</table>", 1);
            case "col":
              return s("<table><colgroup>" + Gr + "</colgroup></table>", 2);
            case "tr":
              return s("<table><tbody>" + Gr + "</tbody></table>", 2);
            case "td":
            case "th":
              return s("<table><tbody><tr>" + Gr + "</tr></tbody></table>", 3);
            case "script":
            case "style":
              return s("<div>" + Gr + "</div>", 1);
            default:
              return s(Gr, 0);
          }
        }
        function ie(ze) {
          ze && ze();
        }
        function I(ze, Qr) {
          return Object.prototype.toString.call(ze) === "[object " + Qr + "]";
        }
        function k(ze) {
          return I(ze, "Function");
        }
        function P(ze) {
          return I(ze, "Object");
        }
        function ae(ze) {
          var Qr = "htmx-internal-data", Wr = ze[Qr];
          return Wr || (Wr = ze[Qr] = {}), Wr;
        }
        function M(ze) {
          var Qr = [];
          if (ze)
            for (var Wr = 0; Wr < ze.length; Wr++)
              Qr.push(ze[Wr]);
          return Qr;
        }
        function oe(ze, Qr) {
          if (ze)
            for (var Wr = 0; Wr < ze.length; Wr++)
              Qr(ze[Wr]);
        }
        function X(ze) {
          var Qr = ze.getBoundingClientRect(), Wr = Qr.top, Gr = Qr.bottom;
          return Wr < window.innerHeight && Gr >= 0;
        }
        function se(ze) {
          return ze.getRootNode && ze.getRootNode() instanceof window.ShadowRoot ? re().body.contains(ze.getRootNode().host) : re().body.contains(ze);
        }
        function D(ze) {
          return ze.trim().split(/\s+/);
        }
        function le(ze, Qr) {
          for (var Wr in Qr)
            Qr.hasOwnProperty(Wr) && (ze[Wr] = Qr[Wr]);
          return ze;
        }
        function E(ze) {
          try {
            return JSON.parse(ze);
          } catch (Qr) {
            return b(Qr), null;
          }
        }
        function U() {
          var ze = "htmx:localStorageTest";
          try {
            return localStorage.setItem(ze, ze), localStorage.removeItem(ze), !0;
          } catch {
            return !1;
          }
        }
        function B(ze) {
          try {
            var Qr = new URL(ze);
            return Qr && (ze = Qr.pathname + Qr.search), /^\/$/.test(ze) || (ze = ze.replace(/\/+$/, "")), ze;
          } catch {
            return ze;
          }
        }
        function t(e) {
          return Tr(re().body, function() {
            return eval(e);
          });
        }
        function F(ze) {
          var Qr = Q.on("htmx:load", function(Wr) {
            ze(Wr.detail.elt);
          });
          return Qr;
        }
        function V() {
          Q.logger = function(ze, Qr, Wr) {
            console && console.log(Qr, ze, Wr);
          };
        }
        function j() {
          Q.logger = null;
        }
        function C(ze, Qr) {
          return Qr ? ze.querySelector(Qr) : C(re(), ze);
        }
        function f(ze, Qr) {
          return Qr ? ze.querySelectorAll(Qr) : f(re(), ze);
        }
        function _(ze, Qr) {
          ze = p(ze), Qr ? setTimeout(function() {
            _(ze), ze = null;
          }, Qr) : ze.parentElement.removeChild(ze);
        }
        function z(ze, Qr, Wr) {
          ze = p(ze), Wr ? setTimeout(function() {
            z(ze, Qr), ze = null;
          }, Wr) : ze.classList && ze.classList.add(Qr);
        }
        function n(ze, Qr, Wr) {
          ze = p(ze), Wr ? setTimeout(function() {
            n(ze, Qr), ze = null;
          }, Wr) : ze.classList && (ze.classList.remove(Qr), ze.classList.length === 0 && ze.removeAttribute("class"));
        }
        function $(ze, Qr) {
          ze = p(ze), ze.classList.toggle(Qr);
        }
        function W(ze, Qr) {
          ze = p(ze), oe(ze.parentElement.children, function(Wr) {
            n(Wr, Qr);
          }), z(ze, Qr);
        }
        function v(ze, Qr) {
          if (ze = p(ze), ze.closest)
            return ze.closest(Qr);
          do
            if (ze == null || h(ze, Qr))
              return ze;
          while (ze = ze && u(ze));
          return null;
        }
        function g(ze, Qr) {
          return ze.substring(0, Qr.length) === Qr;
        }
        function G(ze, Qr) {
          return ze.substring(ze.length - Qr.length) === Qr;
        }
        function J(ze) {
          var Qr = ze.trim();
          return g(Qr, "<") && G(Qr, "/>") ? Qr.substring(1, Qr.length - 2) : Qr;
        }
        function Z(ze, Qr) {
          return Qr.indexOf("closest ") === 0 ? [v(ze, J(Qr.substr(8)))] : Qr.indexOf("find ") === 0 ? [C(ze, J(Qr.substr(5)))] : Qr === "next" ? [ze.nextElementSibling] : Qr.indexOf("next ") === 0 ? [K(ze, J(Qr.substr(5)))] : Qr === "previous" ? [ze.previousElementSibling] : Qr.indexOf("previous ") === 0 ? [Y(ze, J(Qr.substr(9)))] : Qr === "document" ? [document] : Qr === "window" ? [window] : Qr === "body" ? [document.body] : re().querySelectorAll(J(Qr));
        }
        var K = function(ze, Qr) {
          for (var Wr = re().querySelectorAll(Qr), Gr = 0; Gr < Wr.length; Gr++) {
            var Yr = Wr[Gr];
            if (Yr.compareDocumentPosition(ze) === Node.DOCUMENT_POSITION_PRECEDING)
              return Yr;
          }
        }, Y = function(ze, Qr) {
          for (var Wr = re().querySelectorAll(Qr), Gr = Wr.length - 1; Gr >= 0; Gr--) {
            var Yr = Wr[Gr];
            if (Yr.compareDocumentPosition(ze) === Node.DOCUMENT_POSITION_FOLLOWING)
              return Yr;
          }
        };
        function ue(ze, Qr) {
          return Qr ? Z(ze, Qr)[0] : Z(re().body, ze)[0];
        }
        function p(ze) {
          return I(ze, "String") ? C(ze) : ze;
        }
        function ve(ze, Qr, Wr) {
          return k(Qr) ? { target: re().body, event: ze, listener: Qr } : { target: p(ze), event: Qr, listener: Wr };
        }
        function de(ze, Qr, Wr) {
          jr(function() {
            var Yr = ve(ze, Qr, Wr);
            Yr.target.addEventListener(Yr.event, Yr.listener);
          });
          var Gr = k(Qr);
          return Gr ? Qr : Wr;
        }
        function ge(ze, Qr, Wr) {
          return jr(function() {
            var Gr = ve(ze, Qr, Wr);
            Gr.target.removeEventListener(Gr.event, Gr.listener);
          }), k(Qr) ? Qr : Wr;
        }
        var pe = re().createElement("output");
        function me(ze, Qr) {
          var Wr = ne(ze, Qr);
          if (Wr) {
            if (Wr === "this")
              return [xe(ze, Qr)];
            var Gr = Z(ze, Wr);
            return Gr.length === 0 ? (b('The selector "' + Wr + '" on ' + Qr + " returned no matches!"), [pe]) : Gr;
          }
        }
        function xe(ze, Qr) {
          return c(ze, function(Wr) {
            return te(Wr, Qr) != null;
          });
        }
        function ye(ze) {
          var Qr = ne(ze, "hx-target");
          if (Qr)
            return Qr === "this" ? xe(ze, "hx-target") : ue(ze, Qr);
          var Wr = ae(ze);
          return Wr.boosted ? re().body : ze;
        }
        function be(ze) {
          for (var Qr = Q.config.attributesToSettle, Wr = 0; Wr < Qr.length; Wr++)
            if (ze === Qr[Wr])
              return !0;
          return !1;
        }
        function we(ze, Qr) {
          oe(ze.attributes, function(Wr) {
            !Qr.hasAttribute(Wr.name) && be(Wr.name) && ze.removeAttribute(Wr.name);
          }), oe(Qr.attributes, function(Wr) {
            be(Wr.name) && ze.setAttribute(Wr.name, Wr.value);
          });
        }
        function Se(ze, Qr) {
          for (var Wr = Fr(Qr), Gr = 0; Gr < Wr.length; Gr++) {
            var Yr = Wr[Gr];
            try {
              if (Yr.isInlineSwap(ze))
                return !0;
            } catch (Jr) {
              b(Jr);
            }
          }
          return ze === "outerHTML";
        }
        function Ee(ze, Qr, Wr) {
          var Gr = "#" + ee(Qr, "id"), Yr = "outerHTML";
          ze === "true" || (ze.indexOf(":") > 0 ? (Yr = ze.substr(0, ze.indexOf(":")), Gr = ze.substr(ze.indexOf(":") + 1, ze.length)) : Yr = ze);
          var Jr = re().querySelectorAll(Gr);
          return Jr ? (oe(Jr, function(Zr) {
            var Kr, en = Qr.cloneNode(!0);
            Kr = re().createDocumentFragment(), Kr.appendChild(en), Se(Yr, Zr) || (Kr = en);
            var tn = { shouldSwap: !0, target: Zr, fragment: Kr };
            ce(Zr, "htmx:oobBeforeSwap", tn) && (Zr = tn.target, tn.shouldSwap && Fe(Yr, Zr, Zr, Kr, Wr), oe(Wr.elts, function(nn) {
              ce(nn, "htmx:oobAfterSwap", tn);
            }));
          }), Qr.parentNode.removeChild(Qr)) : (Qr.parentNode.removeChild(Qr), fe(re().body, "htmx:oobErrorNoTarget", { content: Qr })), ze;
        }
        function Ce(ze, Qr, Wr) {
          var Gr = ne(ze, "hx-select-oob");
          if (Gr)
            for (var Yr = Gr.split(","), Jr = 0; Jr < Yr.length; Jr++) {
              var Zr = Yr[Jr].split(":", 2), Kr = Zr[0].trim();
              Kr.indexOf("#") === 0 && (Kr = Kr.substring(1));
              var en = Zr[1] || "true", tn = Qr.querySelector("#" + Kr);
              tn && Ee(en, tn, Wr);
            }
          oe(f(Qr, "[hx-swap-oob], [data-hx-swap-oob]"), function(nn) {
            var rn = te(nn, "hx-swap-oob");
            rn != null && Ee(rn, nn, Wr);
          });
        }
        function Re(ze) {
          oe(f(ze, "[hx-preserve], [data-hx-preserve]"), function(Qr) {
            var Wr = te(Qr, "id"), Gr = re().getElementById(Wr);
            Gr != null && Qr.parentNode.replaceChild(Gr, Qr);
          });
        }
        function Te(ze, Qr, Wr) {
          oe(Qr.querySelectorAll("[id]"), function(Gr) {
            var Yr = ee(Gr, "id");
            if (Yr && Yr.length > 0) {
              var Jr = Yr.replace("'", "\\'"), Zr = Gr.tagName.replace(":", "\\:"), Kr = ze.querySelector(Zr + "[id='" + Jr + "']");
              if (Kr && Kr !== ze) {
                var en = Gr.cloneNode();
                we(Gr, Kr), Wr.tasks.push(function() {
                  we(Gr, en);
                });
              }
            }
          });
        }
        function Oe(ze) {
          return function() {
            n(ze, Q.config.addedClass), zt(ze), Nt(ze), qe(ze), ce(ze, "htmx:load");
          };
        }
        function qe(ze) {
          var Qr = "[autofocus]", Wr = h(ze, Qr) ? ze : ze.querySelector(Qr);
          Wr != null && Wr.focus();
        }
        function a(ze, Qr, Wr, Gr) {
          for (Te(ze, Wr, Gr); Wr.childNodes.length > 0; ) {
            var Yr = Wr.firstChild;
            z(Yr, Q.config.addedClass), ze.insertBefore(Yr, Qr), Yr.nodeType !== Node.TEXT_NODE && Yr.nodeType !== Node.COMMENT_NODE && Gr.tasks.push(Oe(Yr));
          }
        }
        function He(ze, Qr) {
          for (var Wr = 0; Wr < ze.length; )
            Qr = (Qr << 5) - Qr + ze.charCodeAt(Wr++) | 0;
          return Qr;
        }
        function Le(ze) {
          var Qr = 0;
          if (ze.attributes)
            for (var Wr = 0; Wr < ze.attributes.length; Wr++) {
              var Gr = ze.attributes[Wr];
              Gr.value && (Qr = He(Gr.name, Qr), Qr = He(Gr.value, Qr));
            }
          return Qr;
        }
        function Ae(ze) {
          var Qr = ae(ze);
          if (Qr.onHandlers) {
            for (var Wr = 0; Wr < Qr.onHandlers.length; Wr++) {
              const Gr = Qr.onHandlers[Wr];
              ze.removeEventListener(Gr.event, Gr.listener);
            }
            delete Qr.onHandlers;
          }
        }
        function Ne(ze) {
          var Qr = ae(ze);
          Qr.timeout && clearTimeout(Qr.timeout), Qr.webSocket && Qr.webSocket.close(), Qr.sseEventSource && Qr.sseEventSource.close(), Qr.listenerInfos && oe(Qr.listenerInfos, function(Wr) {
            Wr.on && Wr.on.removeEventListener(Wr.trigger, Wr.listener);
          }), Ae(ze), oe(Object.keys(Qr), function(Wr) {
            delete Qr[Wr];
          });
        }
        function m(ze) {
          ce(ze, "htmx:beforeCleanupElement"), Ne(ze), ze.children && oe(ze.children, function(Qr) {
            m(Qr);
          });
        }
        function Ie(ze, Qr, Wr) {
          if (ze.tagName === "BODY")
            return Ue(ze, Qr, Wr);
          var Gr, Yr = ze.previousSibling;
          for (a(u(ze), ze, Qr, Wr), Yr == null ? Gr = u(ze).firstChild : Gr = Yr.nextSibling, Wr.elts = Wr.elts.filter(function(Jr) {
            return Jr != ze;
          }); Gr && Gr !== ze; )
            Gr.nodeType === Node.ELEMENT_NODE && Wr.elts.push(Gr), Gr = Gr.nextElementSibling;
          m(ze), u(ze).removeChild(ze);
        }
        function ke(ze, Qr, Wr) {
          return a(ze, ze.firstChild, Qr, Wr);
        }
        function Pe(ze, Qr, Wr) {
          return a(u(ze), ze, Qr, Wr);
        }
        function Me(ze, Qr, Wr) {
          return a(ze, null, Qr, Wr);
        }
        function Xe(ze, Qr, Wr) {
          return a(u(ze), ze.nextSibling, Qr, Wr);
        }
        function De(ze, Qr, Wr) {
          return m(ze), u(ze).removeChild(ze);
        }
        function Ue(ze, Qr, Wr) {
          var Gr = ze.firstChild;
          if (a(ze, Gr, Qr, Wr), Gr) {
            for (; Gr.nextSibling; )
              m(Gr.nextSibling), ze.removeChild(Gr.nextSibling);
            m(Gr), ze.removeChild(Gr);
          }
        }
        function Be(ze, Qr, Wr) {
          var Gr = Wr || ne(ze, "hx-select");
          if (Gr) {
            var Yr = re().createDocumentFragment();
            oe(Qr.querySelectorAll(Gr), function(Jr) {
              Yr.appendChild(Jr);
            }), Qr = Yr;
          }
          return Qr;
        }
        function Fe(ze, Qr, Wr, Gr, Yr) {
          switch (ze) {
            case "none":
              return;
            case "outerHTML":
              Ie(Wr, Gr, Yr);
              return;
            case "afterbegin":
              ke(Wr, Gr, Yr);
              return;
            case "beforebegin":
              Pe(Wr, Gr, Yr);
              return;
            case "beforeend":
              Me(Wr, Gr, Yr);
              return;
            case "afterend":
              Xe(Wr, Gr, Yr);
              return;
            case "delete":
              De(Wr);
              return;
            default:
              for (var Jr = Fr(Qr), Zr = 0; Zr < Jr.length; Zr++) {
                var Kr = Jr[Zr];
                try {
                  var en = Kr.handleSwap(ze, Wr, Gr, Yr);
                  if (en) {
                    if (typeof en.length < "u")
                      for (var tn = 0; tn < en.length; tn++) {
                        var nn = en[tn];
                        nn.nodeType !== Node.TEXT_NODE && nn.nodeType !== Node.COMMENT_NODE && Yr.tasks.push(Oe(nn));
                      }
                    return;
                  }
                } catch (rn) {
                  b(rn);
                }
              }
              ze === "innerHTML" ? Ue(Wr, Gr, Yr) : Fe(Q.config.defaultSwapStyle, Qr, Wr, Gr, Yr);
          }
        }
        function Ve(ze) {
          if (ze.indexOf("<title") > -1) {
            var Qr = ze.replace(H, ""), Wr = Qr.match(q);
            if (Wr)
              return Wr[2];
          }
        }
        function je(ze, Qr, Wr, Gr, Yr, Jr) {
          Yr.title = Ve(Gr);
          var Zr = l(Gr);
          if (Zr)
            return Ce(Wr, Zr, Yr), Zr = Be(Wr, Zr, Jr), Re(Zr), Fe(ze, Wr, Qr, Zr, Yr);
        }
        function _e(ze, Qr, Wr) {
          var Gr = ze.getResponseHeader(Qr);
          if (Gr.indexOf("{") === 0) {
            var Yr = E(Gr);
            for (var Jr in Yr)
              if (Yr.hasOwnProperty(Jr)) {
                var Zr = Yr[Jr];
                P(Zr) || (Zr = { value: Zr }), ce(Wr, Jr, Zr);
              }
          } else
            for (var Kr = Gr.split(","), en = 0; en < Kr.length; en++)
              ce(Wr, Kr[en].trim(), []);
        }
        var x = /[\s,]/, $e = /[_$a-zA-Z]/, We = /[_$a-zA-Z0-9]/, Ge = ['"', "'", "/"], Je = /[^\s]/, Ze = /[{(]/, Ke = /[})]/;
        function Ye(ze) {
          for (var Qr = [], Wr = 0; Wr < ze.length; ) {
            if ($e.exec(ze.charAt(Wr))) {
              for (var Gr = Wr; We.exec(ze.charAt(Wr + 1)); )
                Wr++;
              Qr.push(ze.substr(Gr, Wr - Gr + 1));
            } else if (Ge.indexOf(ze.charAt(Wr)) !== -1) {
              var Yr = ze.charAt(Wr), Gr = Wr;
              for (Wr++; Wr < ze.length && ze.charAt(Wr) !== Yr; )
                ze.charAt(Wr) === "\\" && Wr++, Wr++;
              Qr.push(ze.substr(Gr, Wr - Gr + 1));
            } else {
              var Jr = ze.charAt(Wr);
              Qr.push(Jr);
            }
            Wr++;
          }
          return Qr;
        }
        function Qe(ze, Qr, Wr) {
          return $e.exec(ze.charAt(0)) && ze !== "true" && ze !== "false" && ze !== "this" && ze !== Wr && Qr !== ".";
        }
        function et(ze, Qr, Wr) {
          if (Qr[0] === "[") {
            Qr.shift();
            for (var Gr = 1, Yr = " return (function(" + Wr + "){ return (", Jr = null; Qr.length > 0; ) {
              var Zr = Qr[0];
              if (Zr === "]") {
                if (Gr--, Gr === 0) {
                  Jr === null && (Yr = Yr + "true"), Qr.shift(), Yr += ")})";
                  try {
                    var Kr = Tr(ze, function() {
                      return Function(Yr)();
                    }, function() {
                      return !0;
                    });
                    return Kr.source = Yr, Kr;
                  } catch (en) {
                    return fe(re().body, "htmx:syntax:error", { error: en, source: Yr }), null;
                  }
                }
              } else Zr === "[" && Gr++;
              Qe(Zr, Jr, Wr) ? Yr += "((" + Wr + "." + Zr + ") ? (" + Wr + "." + Zr + ") : (window." + Zr + "))" : Yr = Yr + Zr, Jr = Qr.shift();
            }
          }
        }
        function y(ze, Qr) {
          for (var Wr = ""; ze.length > 0 && !Qr.test(ze[0]); )
            Wr += ze.shift();
          return Wr;
        }
        function tt(ze) {
          var Qr;
          return ze.length > 0 && Ze.test(ze[0]) ? (ze.shift(), Qr = y(ze, Ke).trim(), ze.shift()) : Qr = y(ze, x), Qr;
        }
        var rt = "input, textarea, select";
        function nt(ze, Qr, Wr) {
          var Gr = [], Yr = Ye(Qr);
          do {
            y(Yr, Je);
            var Jr = Yr.length, Zr = y(Yr, /[,\[\s]/);
            if (Zr !== "")
              if (Zr === "every") {
                var Kr = { trigger: "every" };
                y(Yr, Je), Kr.pollInterval = d(y(Yr, /[,\[\s]/)), y(Yr, Je);
                var en = et(ze, Yr, "event");
                en && (Kr.eventFilter = en), Gr.push(Kr);
              } else if (Zr.indexOf("sse:") === 0)
                Gr.push({ trigger: "sse", sseEvent: Zr.substr(4) });
              else {
                var tn = { trigger: Zr }, en = et(ze, Yr, "event");
                for (en && (tn.eventFilter = en); Yr.length > 0 && Yr[0] !== ","; ) {
                  y(Yr, Je);
                  var nn = Yr.shift();
                  if (nn === "changed")
                    tn.changed = !0;
                  else if (nn === "once")
                    tn.once = !0;
                  else if (nn === "consume")
                    tn.consume = !0;
                  else if (nn === "delay" && Yr[0] === ":")
                    Yr.shift(), tn.delay = d(y(Yr, x));
                  else if (nn === "from" && Yr[0] === ":") {
                    if (Yr.shift(), Ze.test(Yr[0]))
                      var rn = tt(Yr);
                    else {
                      var rn = y(Yr, x);
                      if (rn === "closest" || rn === "find" || rn === "next" || rn === "previous") {
                        Yr.shift();
                        var an = tt(Yr);
                        an.length > 0 && (rn += " " + an);
                      }
                    }
                    tn.from = rn;
                  } else nn === "target" && Yr[0] === ":" ? (Yr.shift(), tn.target = tt(Yr)) : nn === "throttle" && Yr[0] === ":" ? (Yr.shift(), tn.throttle = d(y(Yr, x))) : nn === "queue" && Yr[0] === ":" ? (Yr.shift(), tn.queue = y(Yr, x)) : nn === "root" && Yr[0] === ":" ? (Yr.shift(), tn[nn] = tt(Yr)) : nn === "threshold" && Yr[0] === ":" ? (Yr.shift(), tn[nn] = y(Yr, x)) : fe(ze, "htmx:syntax:error", { token: Yr.shift() });
                }
                Gr.push(tn);
              }
            Yr.length === Jr && fe(ze, "htmx:syntax:error", { token: Yr.shift() }), y(Yr, Je);
          } while (Yr[0] === "," && Yr.shift());
          return Wr && (Wr[Qr] = Gr), Gr;
        }
        function it(ze) {
          var Qr = te(ze, "hx-trigger"), Wr = [];
          if (Qr) {
            var Gr = Q.config.triggerSpecsCache;
            Wr = Gr && Gr[Qr] || nt(ze, Qr, Gr);
          }
          return Wr.length > 0 ? Wr : h(ze, "form") ? [{ trigger: "submit" }] : h(ze, 'input[type="button"], input[type="submit"]') ? [{ trigger: "click" }] : h(ze, rt) ? [{ trigger: "change" }] : [{ trigger: "click" }];
        }
        function at(ze) {
          ae(ze).cancelled = !0;
        }
        function ot(ze, Qr, Wr) {
          var Gr = ae(ze);
          Gr.timeout = setTimeout(function() {
            se(ze) && Gr.cancelled !== !0 && (ct(Wr, ze, Wt("hx:poll:trigger", { triggerSpec: Wr, target: ze })) || Qr(ze), ot(ze, Qr, Wr));
          }, Wr.pollInterval);
        }
        function st(ze) {
          return location.hostname === ze.hostname && ee(ze, "href") && ee(ze, "href").indexOf("#") !== 0;
        }
        function lt(ze, Qr, Wr) {
          if (ze.tagName === "A" && st(ze) && (ze.target === "" || ze.target === "_self") || ze.tagName === "FORM") {
            Qr.boosted = !0;
            var Gr, Yr;
            if (ze.tagName === "A")
              Gr = "get", Yr = ee(ze, "href");
            else {
              var Jr = ee(ze, "method");
              Gr = Jr ? Jr.toLowerCase() : "get", Yr = ee(ze, "action");
            }
            Wr.forEach(function(Zr) {
              ht(ze, function(Kr, en) {
                if (v(Kr, Q.config.disableSelector)) {
                  m(Kr);
                  return;
                }
                he(Gr, Yr, Kr, en);
              }, Qr, Zr, !0);
            });
          }
        }
        function ut(ze, Qr) {
          return !!((ze.type === "submit" || ze.type === "click") && (Qr.tagName === "FORM" || h(Qr, 'input[type="submit"], button') && v(Qr, "form") !== null || Qr.tagName === "A" && Qr.href && (Qr.getAttribute("href") === "#" || Qr.getAttribute("href").indexOf("#") !== 0)));
        }
        function ft(ze, Qr) {
          return ae(ze).boosted && ze.tagName === "A" && Qr.type === "click" && (Qr.ctrlKey || Qr.metaKey);
        }
        function ct(ze, Qr, Wr) {
          var Gr = ze.eventFilter;
          if (Gr)
            try {
              return Gr.call(Qr, Wr) !== !0;
            } catch (Yr) {
              return fe(re().body, "htmx:eventFilter:error", { error: Yr, source: Gr.source }), !0;
            }
          return !1;
        }
        function ht(ze, Qr, Wr, Gr, Yr) {
          var Jr = ae(ze), Zr;
          Gr.from ? Zr = Z(ze, Gr.from) : Zr = [ze], Gr.changed && Zr.forEach(function(Kr) {
            var en = ae(Kr);
            en.lastValue = Kr.value;
          }), oe(Zr, function(Kr) {
            var en = function(tn) {
              if (!se(ze)) {
                Kr.removeEventListener(Gr.trigger, en);
                return;
              }
              if (!ft(ze, tn) && ((Yr || ut(tn, ze)) && tn.preventDefault(), !ct(Gr, ze, tn))) {
                var nn = ae(tn);
                if (nn.triggerSpec = Gr, nn.handledFor == null && (nn.handledFor = []), nn.handledFor.indexOf(ze) < 0) {
                  if (nn.handledFor.push(ze), Gr.consume && tn.stopPropagation(), Gr.target && tn.target && !h(tn.target, Gr.target))
                    return;
                  if (Gr.once) {
                    if (Jr.triggeredOnce)
                      return;
                    Jr.triggeredOnce = !0;
                  }
                  if (Gr.changed) {
                    var rn = ae(Kr);
                    if (rn.lastValue === Kr.value)
                      return;
                    rn.lastValue = Kr.value;
                  }
                  if (Jr.delayed && clearTimeout(Jr.delayed), Jr.throttle)
                    return;
                  Gr.throttle > 0 ? Jr.throttle || (Qr(ze, tn), Jr.throttle = setTimeout(function() {
                    Jr.throttle = null;
                  }, Gr.throttle)) : Gr.delay > 0 ? Jr.delayed = setTimeout(function() {
                    Qr(ze, tn);
                  }, Gr.delay) : (ce(ze, "htmx:trigger"), Qr(ze, tn));
                }
              }
            };
            Wr.listenerInfos == null && (Wr.listenerInfos = []), Wr.listenerInfos.push({ trigger: Gr.trigger, listener: en, on: Kr }), Kr.addEventListener(Gr.trigger, en);
          });
        }
        var vt = !1, dt = null;
        function gt() {
          dt || (dt = function() {
            vt = !0;
          }, window.addEventListener("scroll", dt), setInterval(function() {
            vt && (vt = !1, oe(re().querySelectorAll("[hx-trigger='revealed'],[data-hx-trigger='revealed']"), function(ze) {
              pt(ze);
            }));
          }, 200));
        }
        function pt(ze) {
          if (!o(ze, "data-hx-revealed") && X(ze)) {
            ze.setAttribute("data-hx-revealed", "true");
            var Qr = ae(ze);
            Qr.initHash ? ce(ze, "revealed") : ze.addEventListener("htmx:afterProcessNode", function(Wr) {
              ce(ze, "revealed");
            }, { once: !0 });
          }
        }
        function mt(ze, Qr, Wr) {
          for (var Gr = D(Wr), Yr = 0; Yr < Gr.length; Yr++) {
            var Jr = Gr[Yr].split(/:(.+)/);
            Jr[0] === "connect" && xt(ze, Jr[1], 0), Jr[0] === "send" && bt(ze);
          }
        }
        function xt(ze, Qr, Wr) {
          if (se(ze)) {
            if (Qr.indexOf("/") == 0) {
              var Gr = location.hostname + (location.port ? ":" + location.port : "");
              location.protocol == "https:" ? Qr = "wss://" + Gr + Qr : location.protocol == "http:" && (Qr = "ws://" + Gr + Qr);
            }
            var Yr = Q.createWebSocket(Qr);
            Yr.onerror = function(Jr) {
              fe(ze, "htmx:wsError", { error: Jr, socket: Yr }), yt(ze);
            }, Yr.onclose = function(Jr) {
              if ([1006, 1012, 1013].indexOf(Jr.code) >= 0) {
                var Zr = wt(Wr);
                setTimeout(function() {
                  xt(ze, Qr, Wr + 1);
                }, Zr);
              }
            }, Yr.onopen = function(Jr) {
              Wr = 0;
            }, ae(ze).webSocket = Yr, Yr.addEventListener("message", function(Jr) {
              if (!yt(ze)) {
                var Zr = Jr.data;
                R(ze, function(an) {
                  Zr = an.transformResponse(Zr, null, ze);
                });
                for (var Kr = T(ze), en = l(Zr), tn = M(en.children), nn = 0; nn < tn.length; nn++) {
                  var rn = tn[nn];
                  Ee(te(rn, "hx-swap-oob") || "true", rn, Kr);
                }
                nr(Kr.tasks);
              }
            });
          }
        }
        function yt(ze) {
          if (!se(ze))
            return ae(ze).webSocket.close(), !0;
        }
        function bt(ze) {
          var Qr = c(ze, function(Wr) {
            return ae(Wr).webSocket != null;
          });
          Qr ? ze.addEventListener(it(ze)[0].trigger, function(Wr) {
            var Gr = ae(Qr).webSocket, Yr = xr(ze, Qr), Jr = dr(ze, "post"), Zr = Jr.errors, Kr = Jr.values, en = Hr(ze), tn = le(Kr, en), nn = yr(tn, ze);
            if (nn.HEADERS = Yr, Zr && Zr.length > 0) {
              ce(ze, "htmx:validation:halted", Zr);
              return;
            }
            Gr.send(JSON.stringify(nn)), ut(Wr, ze) && Wr.preventDefault();
          }) : fe(ze, "htmx:noWebSocketSourceError");
        }
        function wt(ze) {
          var Qr = Q.config.wsReconnectDelay;
          if (typeof Qr == "function")
            return Qr(ze);
          if (Qr === "full-jitter") {
            var Wr = Math.min(ze, 6), Gr = 1e3 * Math.pow(2, Wr);
            return Gr * Math.random();
          }
          b('htmx.config.wsReconnectDelay must either be a function or the string "full-jitter"');
        }
        function St(ze, Qr, Wr) {
          for (var Gr = D(Wr), Yr = 0; Yr < Gr.length; Yr++) {
            var Jr = Gr[Yr].split(/:(.+)/);
            Jr[0] === "connect" && Et(ze, Jr[1]), Jr[0] === "swap" && Ct(ze, Jr[1]);
          }
        }
        function Et(ze, Qr) {
          var Wr = Q.createEventSource(Qr);
          Wr.onerror = function(Gr) {
            fe(ze, "htmx:sseError", { error: Gr, source: Wr }), Tt(ze);
          }, ae(ze).sseEventSource = Wr;
        }
        function Ct(ze, Qr) {
          var Wr = c(ze, Ot);
          if (Wr) {
            var Gr = ae(Wr).sseEventSource, Yr = function(Jr) {
              if (!Tt(Wr)) {
                if (!se(ze)) {
                  Gr.removeEventListener(Qr, Yr);
                  return;
                }
                var Zr = Jr.data;
                R(ze, function(nn) {
                  Zr = nn.transformResponse(Zr, null, ze);
                });
                var Kr = wr(ze), en = ye(ze), tn = T(ze);
                je(Kr.swapStyle, en, ze, Zr, tn), nr(tn.tasks), ce(ze, "htmx:sseMessage", Jr);
              }
            };
            ae(ze).sseListener = Yr, Gr.addEventListener(Qr, Yr);
          } else
            fe(ze, "htmx:noSSESourceError");
        }
        function Rt(ze, Qr, Wr) {
          var Gr = c(ze, Ot);
          if (Gr) {
            var Yr = ae(Gr).sseEventSource, Jr = function() {
              Tt(Gr) || (se(ze) ? Qr(ze) : Yr.removeEventListener(Wr, Jr));
            };
            ae(ze).sseListener = Jr, Yr.addEventListener(Wr, Jr);
          } else
            fe(ze, "htmx:noSSESourceError");
        }
        function Tt(ze) {
          if (!se(ze))
            return ae(ze).sseEventSource.close(), !0;
        }
        function Ot(ze) {
          return ae(ze).sseEventSource != null;
        }
        function qt(ze, Qr, Wr, Gr) {
          var Yr = function() {
            Wr.loaded || (Wr.loaded = !0, Qr(ze));
          };
          Gr > 0 ? setTimeout(Yr, Gr) : Yr();
        }
        function Ht(ze, Qr, Wr) {
          var Gr = !1;
          return oe(w, function(Yr) {
            if (o(ze, "hx-" + Yr)) {
              var Jr = te(ze, "hx-" + Yr);
              Gr = !0, Qr.path = Jr, Qr.verb = Yr, Wr.forEach(function(Zr) {
                Lt(ze, Zr, Qr, function(Kr, en) {
                  if (v(Kr, Q.config.disableSelector)) {
                    m(Kr);
                    return;
                  }
                  he(Yr, Jr, Kr, en);
                });
              });
            }
          }), Gr;
        }
        function Lt(ze, Qr, Wr, Gr) {
          if (Qr.sseEvent)
            Rt(ze, Gr, Qr.sseEvent);
          else if (Qr.trigger === "revealed")
            gt(), ht(ze, Gr, Wr, Qr), pt(ze);
          else if (Qr.trigger === "intersect") {
            var Yr = {};
            Qr.root && (Yr.root = ue(ze, Qr.root)), Qr.threshold && (Yr.threshold = parseFloat(Qr.threshold));
            var Jr = new IntersectionObserver(function(Zr) {
              for (var Kr = 0; Kr < Zr.length; Kr++) {
                var en = Zr[Kr];
                if (en.isIntersecting) {
                  ce(ze, "intersect");
                  break;
                }
              }
            }, Yr);
            Jr.observe(ze), ht(ze, Gr, Wr, Qr);
          } else Qr.trigger === "load" ? ct(Qr, ze, Wt("load", { elt: ze })) || qt(ze, Gr, Wr, Qr.delay) : Qr.pollInterval > 0 ? (Wr.polling = !0, ot(ze, Gr, Qr)) : ht(ze, Gr, Wr, Qr);
        }
        function At(ze) {
          if (!ze.htmxExecuted && Q.config.allowScriptTags && (ze.type === "text/javascript" || ze.type === "module" || ze.type === "")) {
            var Qr = re().createElement("script");
            oe(ze.attributes, function(Gr) {
              Qr.setAttribute(Gr.name, Gr.value);
            }), Qr.textContent = ze.textContent, Qr.async = !1, Q.config.inlineScriptNonce && (Qr.nonce = Q.config.inlineScriptNonce);
            var Wr = ze.parentElement;
            try {
              Wr.insertBefore(Qr, ze);
            } catch (Gr) {
              b(Gr);
            } finally {
              ze.parentElement && ze.parentElement.removeChild(ze);
            }
          }
        }
        function Nt(ze) {
          h(ze, "script") && At(ze), oe(f(ze, "script"), function(Qr) {
            At(Qr);
          });
        }
        function It(ze) {
          var Qr = ze.attributes;
          if (!Qr)
            return !1;
          for (var Wr = 0; Wr < Qr.length; Wr++) {
            var Gr = Qr[Wr].name;
            if (g(Gr, "hx-on:") || g(Gr, "data-hx-on:") || g(Gr, "hx-on-") || g(Gr, "data-hx-on-"))
              return !0;
          }
          return !1;
        }
        function kt(ze) {
          var Qr = null, Wr = [];
          if (It(ze) && Wr.push(ze), document.evaluate)
            for (var Gr = document.evaluate('.//*[@*[ starts-with(name(), "hx-on:") or starts-with(name(), "data-hx-on:") or starts-with(name(), "hx-on-") or starts-with(name(), "data-hx-on-") ]]', ze); Qr = Gr.iterateNext(); ) Wr.push(Qr);
          else if (typeof ze.getElementsByTagName == "function")
            for (var Yr = ze.getElementsByTagName("*"), Jr = 0; Jr < Yr.length; Jr++)
              It(Yr[Jr]) && Wr.push(Yr[Jr]);
          return Wr;
        }
        function Pt(ze) {
          if (ze.querySelectorAll) {
            var Qr = ", [hx-boost] a, [data-hx-boost] a, a[hx-boost], a[data-hx-boost]", Wr = ze.querySelectorAll(i + Qr + ", form, [type='submit'], [hx-sse], [data-hx-sse], [hx-ws], [data-hx-ws], [hx-ext], [data-hx-ext], [hx-trigger], [data-hx-trigger], [hx-on], [data-hx-on]");
            return Wr;
          } else
            return [];
        }
        function Mt(ze) {
          var Qr = v(ze.target, "button, input[type='submit']"), Wr = Dt(ze);
          Wr && (Wr.lastButtonClicked = Qr);
        }
        function Xt(ze) {
          var Qr = Dt(ze);
          Qr && (Qr.lastButtonClicked = null);
        }
        function Dt(ze) {
          var Qr = v(ze.target, "button, input[type='submit']");
          if (Qr) {
            var Wr = p("#" + ee(Qr, "form")) || v(Qr, "form");
            if (Wr)
              return ae(Wr);
          }
        }
        function Ut(ze) {
          ze.addEventListener("click", Mt), ze.addEventListener("focusin", Mt), ze.addEventListener("focusout", Xt);
        }
        function Bt(ze) {
          for (var Qr = Ye(ze), Wr = 0, Gr = 0; Gr < Qr.length; Gr++) {
            const Yr = Qr[Gr];
            Yr === "{" ? Wr++ : Yr === "}" && Wr--;
          }
          return Wr;
        }
        function Ft(ze, Qr, Wr) {
          var Gr = ae(ze);
          Array.isArray(Gr.onHandlers) || (Gr.onHandlers = []);
          var Yr, Jr = function(Zr) {
            return Tr(ze, function() {
              Yr || (Yr = new Function("event", Wr)), Yr.call(ze, Zr);
            });
          };
          ze.addEventListener(Qr, Jr), Gr.onHandlers.push({ event: Qr, listener: Jr });
        }
        function Vt(ze) {
          var Qr = te(ze, "hx-on");
          if (Qr) {
            for (var Wr = {}, Gr = Qr.split(`
`), Yr = null, Jr = 0; Gr.length > 0; ) {
              var Zr = Gr.shift(), Kr = Zr.match(/^\s*([a-zA-Z:\-\.]+:)(.*)/);
              Jr === 0 && Kr ? (Zr.split(":"), Yr = Kr[1].slice(0, -1), Wr[Yr] = Kr[2]) : Wr[Yr] += Zr, Jr += Bt(Zr);
            }
            for (var en in Wr)
              Ft(ze, en, Wr[en]);
          }
        }
        function jt(ze) {
          Ae(ze);
          for (var Qr = 0; Qr < ze.attributes.length; Qr++) {
            var Wr = ze.attributes[Qr].name, Gr = ze.attributes[Qr].value;
            if (g(Wr, "hx-on") || g(Wr, "data-hx-on")) {
              var Yr = Wr.indexOf("-on") + 3, Jr = Wr.slice(Yr, Yr + 1);
              if (Jr === "-" || Jr === ":") {
                var Zr = Wr.slice(Yr + 1);
                g(Zr, ":") ? Zr = "htmx" + Zr : g(Zr, "-") ? Zr = "htmx:" + Zr.slice(1) : g(Zr, "htmx-") && (Zr = "htmx:" + Zr.slice(5)), Ft(ze, Zr, Gr);
              }
            }
          }
        }
        function _t(ze) {
          if (v(ze, Q.config.disableSelector)) {
            m(ze);
            return;
          }
          var Qr = ae(ze);
          if (Qr.initHash !== Le(ze)) {
            Ne(ze), Qr.initHash = Le(ze), Vt(ze), ce(ze, "htmx:beforeProcessNode"), ze.value && (Qr.lastValue = ze.value);
            var Wr = it(ze), Gr = Ht(ze, Qr, Wr);
            Gr || (ne(ze, "hx-boost") === "true" ? lt(ze, Qr, Wr) : o(ze, "hx-trigger") && Wr.forEach(function(Zr) {
              Lt(ze, Zr, Qr, function() {
              });
            })), (ze.tagName === "FORM" || ee(ze, "type") === "submit" && o(ze, "form")) && Ut(ze);
            var Yr = te(ze, "hx-sse");
            Yr && St(ze, Qr, Yr);
            var Jr = te(ze, "hx-ws");
            Jr && mt(ze, Qr, Jr), ce(ze, "htmx:afterProcessNode");
          }
        }
        function zt(ze) {
          if (ze = p(ze), v(ze, Q.config.disableSelector)) {
            m(ze);
            return;
          }
          _t(ze), oe(Pt(ze), function(Qr) {
            _t(Qr);
          }), oe(kt(ze), jt);
        }
        function $t(ze) {
          return ze.replace(/([a-z0-9])([A-Z])/g, "$1-$2").toLowerCase();
        }
        function Wt(ze, Qr) {
          var Wr;
          return window.CustomEvent && typeof window.CustomEvent == "function" ? Wr = new CustomEvent(ze, { bubbles: !0, cancelable: !0, detail: Qr }) : (Wr = re().createEvent("CustomEvent"), Wr.initCustomEvent(ze, !0, !0, Qr)), Wr;
        }
        function fe(ze, Qr, Wr) {
          ce(ze, Qr, le({ error: Qr }, Wr));
        }
        function Gt(ze) {
          return ze === "htmx:afterProcessNode";
        }
        function R(ze, Qr) {
          oe(Fr(ze), function(Wr) {
            try {
              Qr(Wr);
            } catch (Gr) {
              b(Gr);
            }
          });
        }
        function b(ze) {
          console.error ? console.error(ze) : console.log && console.log("ERROR: ", ze);
        }
        function ce(ze, Qr, Wr) {
          ze = p(ze), Wr == null && (Wr = {}), Wr.elt = ze;
          var Gr = Wt(Qr, Wr);
          Q.logger && !Gt(Qr) && Q.logger(ze, Qr, Wr), Wr.error && (b(Wr.error), ce(ze, "htmx:error", { errorInfo: Wr }));
          var Yr = ze.dispatchEvent(Gr), Jr = $t(Qr);
          if (Yr && Jr !== Qr) {
            var Zr = Wt(Jr, Gr.detail);
            Yr = Yr && ze.dispatchEvent(Zr);
          }
          return R(ze, function(Kr) {
            Yr = Yr && Kr.onEvent(Qr, Gr) !== !1 && !Gr.defaultPrevented;
          }), Yr;
        }
        var Jt = location.pathname + location.search;
        function Zt() {
          var ze = re().querySelector("[hx-history-elt],[data-hx-history-elt]");
          return ze || re().body;
        }
        function Kt(ze, Qr, Wr, Gr) {
          if (U()) {
            if (Q.config.historyCacheSize <= 0) {
              localStorage.removeItem("htmx-history-cache");
              return;
            }
            ze = B(ze);
            for (var Yr = E(localStorage.getItem("htmx-history-cache")) || [], Jr = 0; Jr < Yr.length; Jr++)
              if (Yr[Jr].url === ze) {
                Yr.splice(Jr, 1);
                break;
              }
            var Zr = { url: ze, content: Qr, title: Wr, scroll: Gr };
            for (ce(re().body, "htmx:historyItemCreated", { item: Zr, cache: Yr }), Yr.push(Zr); Yr.length > Q.config.historyCacheSize; )
              Yr.shift();
            for (; Yr.length > 0; )
              try {
                localStorage.setItem("htmx-history-cache", JSON.stringify(Yr));
                break;
              } catch (Kr) {
                fe(re().body, "htmx:historyCacheError", { cause: Kr, cache: Yr }), Yr.shift();
              }
          }
        }
        function Yt(ze) {
          if (!U())
            return null;
          ze = B(ze);
          for (var Qr = E(localStorage.getItem("htmx-history-cache")) || [], Wr = 0; Wr < Qr.length; Wr++)
            if (Qr[Wr].url === ze)
              return Qr[Wr];
          return null;
        }
        function Qt(ze) {
          var Qr = Q.config.requestClass, Wr = ze.cloneNode(!0);
          return oe(f(Wr, "." + Qr), function(Gr) {
            n(Gr, Qr);
          }), Wr.innerHTML;
        }
        function er() {
          var ze = Zt(), Qr = Jt || location.pathname + location.search, Wr;
          try {
            Wr = re().querySelector('[hx-history="false" i],[data-hx-history="false" i]');
          } catch {
            Wr = re().querySelector('[hx-history="false"],[data-hx-history="false"]');
          }
          Wr || (ce(re().body, "htmx:beforeHistorySave", { path: Qr, historyElt: ze }), Kt(Qr, Qt(ze), re().title, window.scrollY)), Q.config.historyEnabled && history.replaceState({ htmx: !0 }, re().title, window.location.href);
        }
        function tr(ze) {
          Q.config.getCacheBusterParam && (ze = ze.replace(/org\.htmx\.cache-buster=[^&]*&?/, ""), (G(ze, "&") || G(ze, "?")) && (ze = ze.slice(0, -1))), Q.config.historyEnabled && history.pushState({ htmx: !0 }, "", ze), Jt = ze;
        }
        function rr(ze) {
          Q.config.historyEnabled && history.replaceState({ htmx: !0 }, "", ze), Jt = ze;
        }
        function nr(ze) {
          oe(ze, function(Qr) {
            Qr.call();
          });
        }
        function ir(ze) {
          var Qr = new XMLHttpRequest(), Wr = { path: ze, xhr: Qr };
          ce(re().body, "htmx:historyCacheMiss", Wr), Qr.open("GET", ze, !0), Qr.setRequestHeader("HX-Request", "true"), Qr.setRequestHeader("HX-History-Restore-Request", "true"), Qr.setRequestHeader("HX-Current-URL", re().location.href), Qr.onload = function() {
            if (this.status >= 200 && this.status < 400) {
              ce(re().body, "htmx:historyCacheMissLoad", Wr);
              var Gr = l(this.response);
              Gr = Gr.querySelector("[hx-history-elt],[data-hx-history-elt]") || Gr;
              var Yr = Zt(), Jr = T(Yr), Zr = Ve(this.response);
              if (Zr) {
                var Kr = C("title");
                Kr ? Kr.innerHTML = Zr : window.document.title = Zr;
              }
              Ue(Yr, Gr, Jr), nr(Jr.tasks), Jt = ze, ce(re().body, "htmx:historyRestore", { path: ze, cacheMiss: !0, serverResponse: this.response });
            } else
              fe(re().body, "htmx:historyCacheMissLoadError", Wr);
          }, Qr.send();
        }
        function ar(ze) {
          er(), ze = ze || location.pathname + location.search;
          var Qr = Yt(ze);
          if (Qr) {
            var Wr = l(Qr.content), Gr = Zt(), Yr = T(Gr);
            Ue(Gr, Wr, Yr), nr(Yr.tasks), document.title = Qr.title, setTimeout(function() {
              window.scrollTo(0, Qr.scroll);
            }, 0), Jt = ze, ce(re().body, "htmx:historyRestore", { path: ze, item: Qr });
          } else
            Q.config.refreshOnHistoryMiss ? window.location.reload(!0) : ir(ze);
        }
        function or(ze) {
          var Qr = me(ze, "hx-indicator");
          return Qr == null && (Qr = [ze]), oe(Qr, function(Wr) {
            var Gr = ae(Wr);
            Gr.requestCount = (Gr.requestCount || 0) + 1, Wr.classList.add.call(Wr.classList, Q.config.requestClass);
          }), Qr;
        }
        function sr(ze) {
          var Qr = me(ze, "hx-disabled-elt");
          return Qr == null && (Qr = []), oe(Qr, function(Wr) {
            var Gr = ae(Wr);
            Gr.requestCount = (Gr.requestCount || 0) + 1, Wr.setAttribute("disabled", "");
          }), Qr;
        }
        function lr(ze, Qr) {
          oe(ze, function(Wr) {
            var Gr = ae(Wr);
            Gr.requestCount = (Gr.requestCount || 0) - 1, Gr.requestCount === 0 && Wr.classList.remove.call(Wr.classList, Q.config.requestClass);
          }), oe(Qr, function(Wr) {
            var Gr = ae(Wr);
            Gr.requestCount = (Gr.requestCount || 0) - 1, Gr.requestCount === 0 && Wr.removeAttribute("disabled");
          });
        }
        function ur(ze, Qr) {
          for (var Wr = 0; Wr < ze.length; Wr++) {
            var Gr = ze[Wr];
            if (Gr.isSameNode(Qr))
              return !0;
          }
          return !1;
        }
        function fr(ze) {
          return ze.name === "" || ze.name == null || ze.disabled || v(ze, "fieldset[disabled]") || ze.type === "button" || ze.type === "submit" || ze.tagName === "image" || ze.tagName === "reset" || ze.tagName === "file" ? !1 : ze.type === "checkbox" || ze.type === "radio" ? ze.checked : !0;
        }
        function cr(ze, Qr, Wr) {
          if (ze != null && Qr != null) {
            var Gr = Wr[ze];
            Gr === void 0 ? Wr[ze] = Qr : Array.isArray(Gr) ? Array.isArray(Qr) ? Wr[ze] = Gr.concat(Qr) : Gr.push(Qr) : Array.isArray(Qr) ? Wr[ze] = [Gr].concat(Qr) : Wr[ze] = [Gr, Qr];
          }
        }
        function hr(ze, Qr, Wr, Gr, Yr) {
          if (!(Gr == null || ur(ze, Gr))) {
            if (ze.push(Gr), fr(Gr)) {
              var Jr = ee(Gr, "name"), Zr = Gr.value;
              Gr.multiple && Gr.tagName === "SELECT" && (Zr = M(Gr.querySelectorAll("option:checked")).map(function(en) {
                return en.value;
              })), Gr.files && (Zr = M(Gr.files)), cr(Jr, Zr, Qr), Yr && vr(Gr, Wr);
            }
            if (h(Gr, "form")) {
              var Kr = Gr.elements;
              oe(Kr, function(en) {
                hr(ze, Qr, Wr, en, Yr);
              });
            }
          }
        }
        function vr(ze, Qr) {
          ze.willValidate && (ce(ze, "htmx:validation:validate"), ze.checkValidity() || (Qr.push({ elt: ze, message: ze.validationMessage, validity: ze.validity }), ce(ze, "htmx:validation:failed", { message: ze.validationMessage, validity: ze.validity })));
        }
        function dr(ze, Qr) {
          var Wr = [], Gr = {}, Yr = {}, Jr = [], Zr = ae(ze);
          Zr.lastButtonClicked && !se(Zr.lastButtonClicked) && (Zr.lastButtonClicked = null);
          var Kr = h(ze, "form") && ze.noValidate !== !0 || te(ze, "hx-validate") === "true";
          if (Zr.lastButtonClicked && (Kr = Kr && Zr.lastButtonClicked.formNoValidate !== !0), Qr !== "get" && hr(Wr, Yr, Jr, v(ze, "form"), Kr), hr(Wr, Gr, Jr, ze, Kr), Zr.lastButtonClicked || ze.tagName === "BUTTON" || ze.tagName === "INPUT" && ee(ze, "type") === "submit") {
            var en = Zr.lastButtonClicked || ze, tn = ee(en, "name");
            cr(tn, en.value, Yr);
          }
          var nn = me(ze, "hx-include");
          return oe(nn, function(rn) {
            hr(Wr, Gr, Jr, rn, Kr), h(rn, "form") || oe(rn.querySelectorAll(rt), function(an) {
              hr(Wr, Gr, Jr, an, Kr);
            });
          }), Gr = le(Gr, Yr), { errors: Jr, values: Gr };
        }
        function gr(ze, Qr, Wr) {
          ze !== "" && (ze += "&"), String(Wr) === "[object Object]" && (Wr = JSON.stringify(Wr));
          var Gr = encodeURIComponent(Wr);
          return ze += encodeURIComponent(Qr) + "=" + Gr, ze;
        }
        function pr(ze) {
          var Qr = "";
          for (var Wr in ze)
            if (ze.hasOwnProperty(Wr)) {
              var Gr = ze[Wr];
              Array.isArray(Gr) ? oe(Gr, function(Yr) {
                Qr = gr(Qr, Wr, Yr);
              }) : Qr = gr(Qr, Wr, Gr);
            }
          return Qr;
        }
        function mr(ze) {
          var Qr = new FormData();
          for (var Wr in ze)
            if (ze.hasOwnProperty(Wr)) {
              var Gr = ze[Wr];
              Array.isArray(Gr) ? oe(Gr, function(Yr) {
                Qr.append(Wr, Yr);
              }) : Qr.append(Wr, Gr);
            }
          return Qr;
        }
        function xr(ze, Qr, Wr) {
          var Gr = { "HX-Request": "true", "HX-Trigger": ee(ze, "id"), "HX-Trigger-Name": ee(ze, "name"), "HX-Target": te(Qr, "id"), "HX-Current-URL": re().location.href };
          return Rr(ze, "hx-headers", !1, Gr), Wr !== void 0 && (Gr["HX-Prompt"] = Wr), ae(ze).boosted && (Gr["HX-Boosted"] = "true"), Gr;
        }
        function yr(ze, Qr) {
          var Wr = ne(Qr, "hx-params");
          if (Wr) {
            if (Wr === "none")
              return {};
            if (Wr === "*")
              return ze;
            if (Wr.indexOf("not ") === 0)
              return oe(Wr.substr(4).split(","), function(Yr) {
                Yr = Yr.trim(), delete ze[Yr];
              }), ze;
            var Gr = {};
            return oe(Wr.split(","), function(Yr) {
              Yr = Yr.trim(), Gr[Yr] = ze[Yr];
            }), Gr;
          } else
            return ze;
        }
        function br(ze) {
          return ee(ze, "href") && ee(ze, "href").indexOf("#") >= 0;
        }
        function wr(ze, Qr) {
          var Wr = Qr || ne(ze, "hx-swap"), Gr = { swapStyle: ae(ze).boosted ? "innerHTML" : Q.config.defaultSwapStyle, swapDelay: Q.config.defaultSwapDelay, settleDelay: Q.config.defaultSettleDelay };
          if (Q.config.scrollIntoViewOnBoost && ae(ze).boosted && !br(ze) && (Gr.show = "top"), Wr) {
            var Yr = D(Wr);
            if (Yr.length > 0)
              for (var Jr = 0; Jr < Yr.length; Jr++) {
                var Zr = Yr[Jr];
                if (Zr.indexOf("swap:") === 0)
                  Gr.swapDelay = d(Zr.substr(5));
                else if (Zr.indexOf("settle:") === 0)
                  Gr.settleDelay = d(Zr.substr(7));
                else if (Zr.indexOf("transition:") === 0)
                  Gr.transition = Zr.substr(11) === "true";
                else if (Zr.indexOf("ignoreTitle:") === 0)
                  Gr.ignoreTitle = Zr.substr(12) === "true";
                else if (Zr.indexOf("scroll:") === 0) {
                  var Kr = Zr.substr(7), en = Kr.split(":"), tn = en.pop(), nn = en.length > 0 ? en.join(":") : null;
                  Gr.scroll = tn, Gr.scrollTarget = nn;
                } else if (Zr.indexOf("show:") === 0) {
                  var rn = Zr.substr(5), en = rn.split(":"), an = en.pop(), nn = en.length > 0 ? en.join(":") : null;
                  Gr.show = an, Gr.showTarget = nn;
                } else if (Zr.indexOf("focus-scroll:") === 0) {
                  var sn = Zr.substr(13);
                  Gr.focusScroll = sn == "true";
                } else Jr == 0 ? Gr.swapStyle = Zr : b("Unknown modifier in hx-swap: " + Zr);
              }
          }
          return Gr;
        }
        function Sr(ze) {
          return ne(ze, "hx-encoding") === "multipart/form-data" || h(ze, "form") && ee(ze, "enctype") === "multipart/form-data";
        }
        function Er(ze, Qr, Wr) {
          var Gr = null;
          return R(Qr, function(Yr) {
            Gr == null && (Gr = Yr.encodeParameters(ze, Wr, Qr));
          }), Gr ?? (Sr(Qr) ? mr(Wr) : pr(Wr));
        }
        function T(ze) {
          return { tasks: [], elts: [ze] };
        }
        function Cr(ze, Qr) {
          var Wr = ze[0], Gr = ze[ze.length - 1];
          if (Qr.scroll) {
            var Yr = null;
            Qr.scrollTarget && (Yr = ue(Wr, Qr.scrollTarget)), Qr.scroll === "top" && (Wr || Yr) && (Yr = Yr || Wr, Yr.scrollTop = 0), Qr.scroll === "bottom" && (Gr || Yr) && (Yr = Yr || Gr, Yr.scrollTop = Yr.scrollHeight);
          }
          if (Qr.show) {
            var Yr = null;
            if (Qr.showTarget) {
              var Jr = Qr.showTarget;
              Qr.showTarget === "window" && (Jr = "body"), Yr = ue(Wr, Jr);
            }
            Qr.show === "top" && (Wr || Yr) && (Yr = Yr || Wr, Yr.scrollIntoView({ block: "start", behavior: Q.config.scrollBehavior })), Qr.show === "bottom" && (Gr || Yr) && (Yr = Yr || Gr, Yr.scrollIntoView({ block: "end", behavior: Q.config.scrollBehavior }));
          }
        }
        function Rr(ze, Qr, Wr, Gr) {
          if (Gr == null && (Gr = {}), ze == null)
            return Gr;
          var Yr = te(ze, Qr);
          if (Yr) {
            var Jr = Yr.trim(), Zr = Wr;
            if (Jr === "unset")
              return null;
            Jr.indexOf("javascript:") === 0 ? (Jr = Jr.substr(11), Zr = !0) : Jr.indexOf("js:") === 0 && (Jr = Jr.substr(3), Zr = !0), Jr.indexOf("{") !== 0 && (Jr = "{" + Jr + "}");
            var Kr;
            Zr ? Kr = Tr(ze, function() {
              return Function("return (" + Jr + ")")();
            }, {}) : Kr = E(Jr);
            for (var en in Kr)
              Kr.hasOwnProperty(en) && Gr[en] == null && (Gr[en] = Kr[en]);
          }
          return Rr(u(ze), Qr, Wr, Gr);
        }
        function Tr(ze, Qr, Wr) {
          return Q.config.allowEval ? Qr() : (fe(ze, "htmx:evalDisallowedError"), Wr);
        }
        function Or(ze, Qr) {
          return Rr(ze, "hx-vars", !0, Qr);
        }
        function qr(ze, Qr) {
          return Rr(ze, "hx-vals", !1, Qr);
        }
        function Hr(ze) {
          return le(Or(ze), qr(ze));
        }
        function Lr(ze, Qr, Wr) {
          if (Wr !== null)
            try {
              ze.setRequestHeader(Qr, Wr);
            } catch {
              ze.setRequestHeader(Qr, encodeURIComponent(Wr)), ze.setRequestHeader(Qr + "-URI-AutoEncoded", "true");
            }
        }
        function Ar(ze) {
          if (ze.responseURL && typeof URL < "u")
            try {
              var Qr = new URL(ze.responseURL);
              return Qr.pathname + Qr.search;
            } catch {
              fe(re().body, "htmx:badResponseUrl", { url: ze.responseURL });
            }
        }
        function O(ze, Qr) {
          return Qr.test(ze.getAllResponseHeaders());
        }
        function Nr(ze, Qr, Wr) {
          return ze = ze.toLowerCase(), Wr ? Wr instanceof Element || I(Wr, "String") ? he(ze, Qr, null, null, { targetOverride: p(Wr), returnPromise: !0 }) : he(ze, Qr, p(Wr.source), Wr.event, { handler: Wr.handler, headers: Wr.headers, values: Wr.values, targetOverride: p(Wr.target), swapOverride: Wr.swap, select: Wr.select, returnPromise: !0 }) : he(ze, Qr, null, null, { returnPromise: !0 });
        }
        function Ir(ze) {
          for (var Qr = []; ze; )
            Qr.push(ze), ze = ze.parentElement;
          return Qr;
        }
        function kr(ze, Qr, Wr) {
          var Gr, Yr;
          if (typeof URL == "function") {
            Yr = new URL(Qr, document.location.href);
            var Jr = document.location.origin;
            Gr = Jr === Yr.origin;
          } else
            Yr = Qr, Gr = g(Qr, document.location.origin);
          return Q.config.selfRequestsOnly && !Gr ? !1 : ce(ze, "htmx:validateUrl", le({ url: Yr, sameHost: Gr }, Wr));
        }
        function he(ze, Qr, Wr, Gr, Yr, Jr) {
          var Zr = null, Kr = null;
          if (Yr = Yr ?? {}, Yr.returnPromise && typeof Promise < "u")
            var en = new Promise(function(An, Hn) {
              Zr = An, Kr = Hn;
            });
          Wr == null && (Wr = re().body);
          var tn = Yr.handler || Mr, nn = Yr.select || null;
          if (!se(Wr))
            return ie(Zr), en;
          var rn = Yr.targetOverride || ye(Wr);
          if (rn == null || rn == pe)
            return fe(Wr, "htmx:targetError", { target: te(Wr, "hx-target") }), ie(Kr), en;
          var an = ae(Wr), sn = an.lastButtonClicked;
          if (sn) {
            var ln = ee(sn, "formaction");
            ln != null && (Qr = ln);
            var un = ee(sn, "formmethod");
            un != null && un.toLowerCase() !== "dialog" && (ze = un);
          }
          var bn = ne(Wr, "hx-confirm");
          if (Jr === void 0) {
            var Cn = function(An) {
              return he(ze, Qr, Wr, Gr, Yr, !!An);
            }, hn = { target: rn, elt: Wr, path: Qr, verb: ze, triggeringEvent: Gr, etc: Yr, issueRequest: Cn, question: bn };
            if (ce(Wr, "htmx:confirm", hn) === !1)
              return ie(Zr), en;
          }
          var gn = Wr, dn = ne(Wr, "hx-sync"), yn = null, In = !1;
          if (dn) {
            var pn = dn.split(":"), vn = pn[0].trim();
            if (vn === "this" ? gn = xe(Wr, "hx-sync") : gn = ue(Wr, vn), dn = (pn[1] || "drop").trim(), an = ae(gn), dn === "drop" && an.xhr && an.abortable !== !0)
              return ie(Zr), en;
            if (dn === "abort") {
              if (an.xhr)
                return ie(Zr), en;
              In = !0;
            } else if (dn === "replace")
              ce(gn, "htmx:abort");
            else if (dn.indexOf("queue") === 0) {
              var mn = dn.split(" ");
              yn = (mn[1] || "last").trim();
            }
          }
          if (an.xhr)
            if (an.abortable)
              ce(gn, "htmx:abort");
            else {
              if (yn == null) {
                if (Gr) {
                  var En = ae(Gr);
                  En && En.triggerSpec && En.triggerSpec.queue && (yn = En.triggerSpec.queue);
                }
                yn == null && (yn = "last");
              }
              return an.queuedRequests == null && (an.queuedRequests = []), yn === "first" && an.queuedRequests.length === 0 ? an.queuedRequests.push(function() {
                he(ze, Qr, Wr, Gr, Yr);
              }) : yn === "all" ? an.queuedRequests.push(function() {
                he(ze, Qr, Wr, Gr, Yr);
              }) : yn === "last" && (an.queuedRequests = [], an.queuedRequests.push(function() {
                he(ze, Qr, Wr, Gr, Yr);
              })), ie(Zr), en;
            }
          var cn = new XMLHttpRequest();
          an.xhr = cn, an.abortable = In;
          var Sn = function() {
            if (an.xhr = null, an.abortable = !1, an.queuedRequests != null && an.queuedRequests.length > 0) {
              var An = an.queuedRequests.shift();
              An();
            }
          }, xn = ne(Wr, "hx-prompt");
          if (xn) {
            var Dn = prompt(xn);
            if (Dn === null || !ce(Wr, "htmx:prompt", { prompt: Dn, target: rn }))
              return ie(Zr), Sn(), en;
          }
          if (bn && !Jr && !confirm(bn))
            return ie(Zr), Sn(), en;
          var fn = xr(Wr, rn, Dn);
          ze !== "get" && !Sr(Wr) && (fn["Content-Type"] = "application/x-www-form-urlencoded"), Yr.headers && (fn = le(fn, Yr.headers));
          var Tn = dr(Wr, ze), Rn = Tn.errors, wn = Tn.values;
          Yr.values && (wn = le(wn, Yr.values));
          var Xn = Hr(Wr), Bn = le(wn, Xn), Pn = yr(Bn, Wr);
          Q.config.getCacheBusterParam && ze === "get" && (Pn["org.htmx.cache-buster"] = ee(rn, "id") || "true"), (Qr == null || Qr === "") && (Qr = re().location.href);
          var Mn = Rr(Wr, "hx-request"), Un = ae(Wr).boosted, qn = Q.config.methodsThatUseUrlParams.indexOf(ze) >= 0, On = { boosted: Un, useUrlParams: qn, parameters: Pn, unfilteredParameters: Bn, headers: fn, target: rn, verb: ze, errors: Rn, withCredentials: Yr.credentials || Mn.credentials || Q.config.withCredentials, timeout: Yr.timeout || Mn.timeout || Q.config.timeout, path: Qr, triggeringEvent: Gr };
          if (!ce(Wr, "htmx:configRequest", On))
            return ie(Zr), Sn(), en;
          if (Qr = On.path, ze = On.verb, fn = On.headers, Pn = On.parameters, Rn = On.errors, qn = On.useUrlParams, Rn && Rn.length > 0)
            return ce(Wr, "htmx:validation:halted", On), ie(Zr), Sn(), en;
          var Wn = Qr.split("#"), Yn = Wn[0], Qn = Wn[1], Nn = Qr;
          if (qn) {
            Nn = Yn;
            var zn = Object.keys(Pn).length !== 0;
            zn && (Nn.indexOf("?") < 0 ? Nn += "?" : Nn += "&", Nn += pr(Pn), Qn && (Nn += "#" + Qn));
          }
          if (!kr(Wr, Nn, On))
            return fe(Wr, "htmx:invalidPath", On), ie(Kr), en;
          if (cn.open(ze.toUpperCase(), Nn, !0), cn.overrideMimeType("text/html"), cn.withCredentials = On.withCredentials, cn.timeout = On.timeout, !Mn.noHeaders) {
            for (var jn in fn)
              if (fn.hasOwnProperty(jn)) {
                var Vn = fn[jn];
                Lr(cn, jn, Vn);
              }
          }
          var _n = { xhr: cn, target: rn, requestConfig: On, etc: Yr, boosted: Un, select: nn, pathInfo: { requestPath: Qr, finalRequestPath: Nn, anchor: Qn } };
          if (cn.onload = function() {
            try {
              var An = Ir(Wr);
              if (_n.pathInfo.responsePath = Ar(cn), tn(Wr, _n), lr(Fn, kn), ce(Wr, "htmx:afterRequest", _n), ce(Wr, "htmx:afterOnLoad", _n), !se(Wr)) {
                for (var Hn = null; An.length > 0 && Hn == null; ) {
                  var Ln = An.shift();
                  se(Ln) && (Hn = Ln);
                }
                Hn && (ce(Hn, "htmx:afterRequest", _n), ce(Hn, "htmx:afterOnLoad", _n));
              }
              ie(Zr), Sn();
            } catch (Gn) {
              throw fe(Wr, "htmx:onLoadError", le({ error: Gn }, _n)), Gn;
            }
          }, cn.onerror = function() {
            lr(Fn, kn), fe(Wr, "htmx:afterRequest", _n), fe(Wr, "htmx:sendError", _n), ie(Kr), Sn();
          }, cn.onabort = function() {
            lr(Fn, kn), fe(Wr, "htmx:afterRequest", _n), fe(Wr, "htmx:sendAbort", _n), ie(Kr), Sn();
          }, cn.ontimeout = function() {
            lr(Fn, kn), fe(Wr, "htmx:afterRequest", _n), fe(Wr, "htmx:timeout", _n), ie(Kr), Sn();
          }, !ce(Wr, "htmx:beforeRequest", _n))
            return ie(Zr), Sn(), en;
          var Fn = or(Wr), kn = sr(Wr);
          oe(["loadstart", "loadend", "progress", "abort"], function(An) {
            oe([cn, cn.upload], function(Hn) {
              Hn.addEventListener(An, function(Ln) {
                ce(Wr, "htmx:xhr:" + An, { lengthComputable: Ln.lengthComputable, loaded: Ln.loaded, total: Ln.total });
              });
            });
          }), ce(Wr, "htmx:beforeSend", _n);
          var $n = qn ? null : Er(cn, Wr, Pn);
          return cn.send($n), en;
        }
        function Pr(ze, Qr) {
          var Wr = Qr.xhr, Gr = null, Yr = null;
          if (O(Wr, /HX-Push:/i) ? (Gr = Wr.getResponseHeader("HX-Push"), Yr = "push") : O(Wr, /HX-Push-Url:/i) ? (Gr = Wr.getResponseHeader("HX-Push-Url"), Yr = "push") : O(Wr, /HX-Replace-Url:/i) && (Gr = Wr.getResponseHeader("HX-Replace-Url"), Yr = "replace"), Gr)
            return Gr === "false" ? {} : { type: Yr, path: Gr };
          var Jr = Qr.pathInfo.finalRequestPath, Zr = Qr.pathInfo.responsePath, Kr = ne(ze, "hx-push-url"), en = ne(ze, "hx-replace-url"), tn = ae(ze).boosted, nn = null, rn = null;
          return Kr ? (nn = "push", rn = Kr) : en ? (nn = "replace", rn = en) : tn && (nn = "push", rn = Zr || Jr), rn ? rn === "false" ? {} : (rn === "true" && (rn = Zr || Jr), Qr.pathInfo.anchor && rn.indexOf("#") === -1 && (rn = rn + "#" + Qr.pathInfo.anchor), { type: nn, path: rn }) : {};
        }
        function Mr(ze, Qr) {
          var Wr = Qr.xhr, Gr = Qr.target, Yr = Qr.etc;
          Qr.requestConfig;
          var Jr = Qr.select;
          if (ce(ze, "htmx:beforeOnLoad", Qr)) {
            if (O(Wr, /HX-Trigger:/i) && _e(Wr, "HX-Trigger", ze), O(Wr, /HX-Location:/i)) {
              er();
              var Zr = Wr.getResponseHeader("HX-Location"), Kr;
              Zr.indexOf("{") === 0 && (Kr = E(Zr), Zr = Kr.path, delete Kr.path), Nr("GET", Zr, Kr).then(function() {
                tr(Zr);
              });
              return;
            }
            var en = O(Wr, /HX-Refresh:/i) && Wr.getResponseHeader("HX-Refresh") === "true";
            if (O(Wr, /HX-Redirect:/i)) {
              location.href = Wr.getResponseHeader("HX-Redirect"), en && location.reload();
              return;
            }
            if (en) {
              location.reload();
              return;
            }
            O(Wr, /HX-Retarget:/i) && (Wr.getResponseHeader("HX-Retarget") === "this" ? Qr.target = ze : Qr.target = ue(ze, Wr.getResponseHeader("HX-Retarget")));
            var tn = Pr(ze, Qr), nn = Wr.status >= 200 && Wr.status < 400 && Wr.status !== 204, rn = Wr.response, an = Wr.status >= 400, sn = Q.config.ignoreTitle, ln = le({ shouldSwap: nn, serverResponse: rn, isError: an, ignoreTitle: sn }, Qr);
            if (ce(Gr, "htmx:beforeSwap", ln)) {
              if (Gr = ln.target, rn = ln.serverResponse, an = ln.isError, sn = ln.ignoreTitle, Qr.target = Gr, Qr.failed = an, Qr.successful = !an, ln.shouldSwap) {
                Wr.status === 286 && at(ze), R(ze, function(pn) {
                  rn = pn.transformResponse(rn, Wr, ze);
                }), tn.type && er();
                var un = Yr.swapOverride;
                O(Wr, /HX-Reswap:/i) && (un = Wr.getResponseHeader("HX-Reswap"));
                var Kr = wr(ze, un);
                Kr.hasOwnProperty("ignoreTitle") && (sn = Kr.ignoreTitle), Gr.classList.add(Q.config.swappingClass);
                var bn = null, Cn = null, hn = function() {
                  try {
                    var pn = document.activeElement, vn = {};
                    try {
                      vn = { elt: pn, start: pn ? pn.selectionStart : null, end: pn ? pn.selectionEnd : null };
                    } catch {
                    }
                    var mn;
                    Jr && (mn = Jr), O(Wr, /HX-Reselect:/i) && (mn = Wr.getResponseHeader("HX-Reselect")), tn.type && (ce(re().body, "htmx:beforeHistoryUpdate", le({ history: tn }, Qr)), tn.type === "push" ? (tr(tn.path), ce(re().body, "htmx:pushedIntoHistory", { path: tn.path })) : (rr(tn.path), ce(re().body, "htmx:replacedInHistory", { path: tn.path })));
                    var En = T(Gr);
                    if (je(Kr.swapStyle, Gr, ze, rn, En, mn), vn.elt && !se(vn.elt) && ee(vn.elt, "id")) {
                      var cn = document.getElementById(ee(vn.elt, "id")), Sn = { preventScroll: Kr.focusScroll !== void 0 ? !Kr.focusScroll : !Q.config.defaultFocusScroll };
                      if (cn) {
                        if (vn.start && cn.setSelectionRange)
                          try {
                            cn.setSelectionRange(vn.start, vn.end);
                          } catch {
                          }
                        cn.focus(Sn);
                      }
                    }
                    if (Gr.classList.remove(Q.config.swappingClass), oe(En.elts, function(fn) {
                      fn.classList && fn.classList.add(Q.config.settlingClass), ce(fn, "htmx:afterSwap", Qr);
                    }), O(Wr, /HX-Trigger-After-Swap:/i)) {
                      var xn = ze;
                      se(ze) || (xn = re().body), _e(Wr, "HX-Trigger-After-Swap", xn);
                    }
                    var Dn = function() {
                      if (oe(En.tasks, function(wn) {
                        wn.call();
                      }), oe(En.elts, function(wn) {
                        wn.classList && wn.classList.remove(Q.config.settlingClass), ce(wn, "htmx:afterSettle", Qr);
                      }), Qr.pathInfo.anchor) {
                        var fn = re().getElementById(Qr.pathInfo.anchor);
                        fn && fn.scrollIntoView({ block: "start", behavior: "auto" });
                      }
                      if (En.title && !sn) {
                        var Tn = C("title");
                        Tn ? Tn.innerHTML = En.title : window.document.title = En.title;
                      }
                      if (Cr(En.elts, Kr), O(Wr, /HX-Trigger-After-Settle:/i)) {
                        var Rn = ze;
                        se(ze) || (Rn = re().body), _e(Wr, "HX-Trigger-After-Settle", Rn);
                      }
                      ie(bn);
                    };
                    Kr.settleDelay > 0 ? setTimeout(Dn, Kr.settleDelay) : Dn();
                  } catch (fn) {
                    throw fe(ze, "htmx:swapError", Qr), ie(Cn), fn;
                  }
                }, gn = Q.config.globalViewTransitions;
                if (Kr.hasOwnProperty("transition") && (gn = Kr.transition), gn && ce(ze, "htmx:beforeTransition", Qr) && typeof Promise < "u" && document.startViewTransition) {
                  var dn = new Promise(function(pn, vn) {
                    bn = pn, Cn = vn;
                  }), yn = hn;
                  hn = function() {
                    document.startViewTransition(function() {
                      return yn(), dn;
                    });
                  };
                }
                Kr.swapDelay > 0 ? setTimeout(hn, Kr.swapDelay) : hn();
              }
              an && fe(ze, "htmx:responseError", le({ error: "Response Status Error Code " + Wr.status + " from " + Qr.pathInfo.requestPath }, Qr));
            }
          }
        }
        var Xr = {};
        function Dr() {
          return { init: function(ze) {
            return null;
          }, onEvent: function(ze, Qr) {
            return !0;
          }, transformResponse: function(ze, Qr, Wr) {
            return ze;
          }, isInlineSwap: function(ze) {
            return !1;
          }, handleSwap: function(ze, Qr, Wr, Gr) {
            return !1;
          }, encodeParameters: function(ze, Qr, Wr) {
            return null;
          } };
        }
        function Ur(ze, Qr) {
          Qr.init && Qr.init(r), Xr[ze] = le(Dr(), Qr);
        }
        function Br(ze) {
          delete Xr[ze];
        }
        function Fr(ze, Qr, Wr) {
          if (ze == null)
            return Qr;
          Qr == null && (Qr = []), Wr == null && (Wr = []);
          var Gr = te(ze, "hx-ext");
          return Gr && oe(Gr.split(","), function(Yr) {
            if (Yr = Yr.replace(/ /g, ""), Yr.slice(0, 7) == "ignore:") {
              Wr.push(Yr.slice(7));
              return;
            }
            if (Wr.indexOf(Yr) < 0) {
              var Jr = Xr[Yr];
              Jr && Qr.indexOf(Jr) < 0 && Qr.push(Jr);
            }
          }), Fr(u(ze), Qr, Wr);
        }
        var Vr = !1;
        re().addEventListener("DOMContentLoaded", function() {
          Vr = !0;
        });
        function jr(ze) {
          Vr || re().readyState === "complete" ? ze() : re().addEventListener("DOMContentLoaded", ze);
        }
        function _r() {
          Q.config.includeIndicatorStyles !== !1 && re().head.insertAdjacentHTML("beforeend", "<style>                      ." + Q.config.indicatorClass + "{opacity:0}                      ." + Q.config.requestClass + " ." + Q.config.indicatorClass + "{opacity:1; transition: opacity 200ms ease-in;}                      ." + Q.config.requestClass + "." + Q.config.indicatorClass + "{opacity:1; transition: opacity 200ms ease-in;}                    </style>");
        }
        function zr() {
          var ze = re().querySelector('meta[name="htmx-config"]');
          return ze ? E(ze.content) : null;
        }
        function $r() {
          var ze = zr();
          ze && (Q.config = le(Q.config, ze));
        }
        return jr(function() {
          $r(), _r();
          var ze = re().body;
          zt(ze);
          var Qr = re().querySelectorAll("[hx-trigger='restored'],[data-hx-trigger='restored']");
          ze.addEventListener("htmx:abort", function(Gr) {
            var Yr = Gr.target, Jr = ae(Yr);
            Jr && Jr.xhr && Jr.xhr.abort();
          });
          const Wr = window.onpopstate ? window.onpopstate.bind(window) : null;
          window.onpopstate = function(Gr) {
            Gr.state && Gr.state.htmx ? (ar(), oe(Qr, function(Yr) {
              ce(Yr, "htmx:restored", { document: re(), triggerEvent: ce });
            })) : Wr && Wr(Gr);
          }, setTimeout(function() {
            ce(ze, "htmx:load", {}), ze = null;
          }, 0);
        }), Q;
      }();
    });
  }(htmx_min$1)), htmx_min$1.exports;
}
var htmx_minExports = requireHtmx_min();
const htmx = /* @__PURE__ */ getDefaultExportFromCjs(htmx_minExports);
/**!
 * Sortable 1.15.6
 * @author	RubaXa   <trash@rubaxa.org>
 * @author	owenm    <owen23355@gmail.com>
 * @license MIT
 */
function ownKeys(ze, Qr) {
  var Wr = Object.keys(ze);
  if (Object.getOwnPropertySymbols) {
    var Gr = Object.getOwnPropertySymbols(ze);
    Qr && (Gr = Gr.filter(function(Yr) {
      return Object.getOwnPropertyDescriptor(ze, Yr).enumerable;
    })), Wr.push.apply(Wr, Gr);
  }
  return Wr;
}
function _objectSpread2(ze) {
  for (var Qr = 1; Qr < arguments.length; Qr++) {
    var Wr = arguments[Qr] != null ? arguments[Qr] : {};
    Qr % 2 ? ownKeys(Object(Wr), !0).forEach(function(Gr) {
      _defineProperty(ze, Gr, Wr[Gr]);
    }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(ze, Object.getOwnPropertyDescriptors(Wr)) : ownKeys(Object(Wr)).forEach(function(Gr) {
      Object.defineProperty(ze, Gr, Object.getOwnPropertyDescriptor(Wr, Gr));
    });
  }
  return ze;
}
function _typeof(ze) {
  "@babel/helpers - typeof";
  return typeof Symbol == "function" && typeof Symbol.iterator == "symbol" ? _typeof = function(Qr) {
    return typeof Qr;
  } : _typeof = function(Qr) {
    return Qr && typeof Symbol == "function" && Qr.constructor === Symbol && Qr !== Symbol.prototype ? "symbol" : typeof Qr;
  }, _typeof(ze);
}
function _defineProperty(ze, Qr, Wr) {
  return Qr in ze ? Object.defineProperty(ze, Qr, {
    value: Wr,
    enumerable: !0,
    configurable: !0,
    writable: !0
  }) : ze[Qr] = Wr, ze;
}
function _extends() {
  return _extends = Object.assign || function(ze) {
    for (var Qr = 1; Qr < arguments.length; Qr++) {
      var Wr = arguments[Qr];
      for (var Gr in Wr)
        Object.prototype.hasOwnProperty.call(Wr, Gr) && (ze[Gr] = Wr[Gr]);
    }
    return ze;
  }, _extends.apply(this, arguments);
}
function _objectWithoutPropertiesLoose(ze, Qr) {
  if (ze == null) return {};
  var Wr = {}, Gr = Object.keys(ze), Yr, Jr;
  for (Jr = 0; Jr < Gr.length; Jr++)
    Yr = Gr[Jr], !(Qr.indexOf(Yr) >= 0) && (Wr[Yr] = ze[Yr]);
  return Wr;
}
function _objectWithoutProperties(ze, Qr) {
  if (ze == null) return {};
  var Wr = _objectWithoutPropertiesLoose(ze, Qr), Gr, Yr;
  if (Object.getOwnPropertySymbols) {
    var Jr = Object.getOwnPropertySymbols(ze);
    for (Yr = 0; Yr < Jr.length; Yr++)
      Gr = Jr[Yr], !(Qr.indexOf(Gr) >= 0) && Object.prototype.propertyIsEnumerable.call(ze, Gr) && (Wr[Gr] = ze[Gr]);
  }
  return Wr;
}
var version = "1.15.6";
function userAgent(ze) {
  if (typeof window < "u" && window.navigator)
    return !!/* @__PURE__ */ navigator.userAgent.match(ze);
}
var IE11OrLess = userAgent(/(?:Trident.*rv[ :]?11\.|msie|iemobile|Windows Phone)/i), Edge = userAgent(/Edge/i), FireFox = userAgent(/firefox/i), Safari = userAgent(/safari/i) && !userAgent(/chrome/i) && !userAgent(/android/i), IOS = userAgent(/iP(ad|od|hone)/i), ChromeForAndroid = userAgent(/chrome/i) && userAgent(/android/i), captureMode = {
  capture: !1,
  passive: !1
};
function on(ze, Qr, Wr) {
  ze.addEventListener(Qr, Wr, !IE11OrLess && captureMode);
}
function off(ze, Qr, Wr) {
  ze.removeEventListener(Qr, Wr, !IE11OrLess && captureMode);
}
function matches(ze, Qr) {
  if (Qr) {
    if (Qr[0] === ">" && (Qr = Qr.substring(1)), ze)
      try {
        if (ze.matches)
          return ze.matches(Qr);
        if (ze.msMatchesSelector)
          return ze.msMatchesSelector(Qr);
        if (ze.webkitMatchesSelector)
          return ze.webkitMatchesSelector(Qr);
      } catch {
        return !1;
      }
    return !1;
  }
}
function getParentOrHost(ze) {
  return ze.host && ze !== document && ze.host.nodeType ? ze.host : ze.parentNode;
}
function closest(ze, Qr, Wr, Gr) {
  if (ze) {
    Wr = Wr || document;
    do {
      if (Qr != null && (Qr[0] === ">" ? ze.parentNode === Wr && matches(ze, Qr) : matches(ze, Qr)) || Gr && ze === Wr)
        return ze;
      if (ze === Wr) break;
    } while (ze = getParentOrHost(ze));
  }
  return null;
}
var R_SPACE = /\s+/g;
function toggleClass(ze, Qr, Wr) {
  if (ze && Qr)
    if (ze.classList)
      ze.classList[Wr ? "add" : "remove"](Qr);
    else {
      var Gr = (" " + ze.className + " ").replace(R_SPACE, " ").replace(" " + Qr + " ", " ");
      ze.className = (Gr + (Wr ? " " + Qr : "")).replace(R_SPACE, " ");
    }
}
function css(ze, Qr, Wr) {
  var Gr = ze && ze.style;
  if (Gr) {
    if (Wr === void 0)
      return document.defaultView && document.defaultView.getComputedStyle ? Wr = document.defaultView.getComputedStyle(ze, "") : ze.currentStyle && (Wr = ze.currentStyle), Qr === void 0 ? Wr : Wr[Qr];
    !(Qr in Gr) && Qr.indexOf("webkit") === -1 && (Qr = "-webkit-" + Qr), Gr[Qr] = Wr + (typeof Wr == "string" ? "" : "px");
  }
}
function matrix(ze, Qr) {
  var Wr = "";
  if (typeof ze == "string")
    Wr = ze;
  else
    do {
      var Gr = css(ze, "transform");
      Gr && Gr !== "none" && (Wr = Gr + " " + Wr);
    } while (!Qr && (ze = ze.parentNode));
  var Yr = window.DOMMatrix || window.WebKitCSSMatrix || window.CSSMatrix || window.MSCSSMatrix;
  return Yr && new Yr(Wr);
}
function find(ze, Qr, Wr) {
  if (ze) {
    var Gr = ze.getElementsByTagName(Qr), Yr = 0, Jr = Gr.length;
    if (Wr)
      for (; Yr < Jr; Yr++)
        Wr(Gr[Yr], Yr);
    return Gr;
  }
  return [];
}
function getWindowScrollingElement() {
  var ze = document.scrollingElement;
  return ze || document.documentElement;
}
function getRect(ze, Qr, Wr, Gr, Yr) {
  if (!(!ze.getBoundingClientRect && ze !== window)) {
    var Jr, Zr, Kr, en, tn, nn, rn;
    if (ze !== window && ze.parentNode && ze !== getWindowScrollingElement() ? (Jr = ze.getBoundingClientRect(), Zr = Jr.top, Kr = Jr.left, en = Jr.bottom, tn = Jr.right, nn = Jr.height, rn = Jr.width) : (Zr = 0, Kr = 0, en = window.innerHeight, tn = window.innerWidth, nn = window.innerHeight, rn = window.innerWidth), (Qr || Wr) && ze !== window && (Yr = Yr || ze.parentNode, !IE11OrLess))
      do
        if (Yr && Yr.getBoundingClientRect && (css(Yr, "transform") !== "none" || Wr && css(Yr, "position") !== "static")) {
          var an = Yr.getBoundingClientRect();
          Zr -= an.top + parseInt(css(Yr, "border-top-width")), Kr -= an.left + parseInt(css(Yr, "border-left-width")), en = Zr + Jr.height, tn = Kr + Jr.width;
          break;
        }
      while (Yr = Yr.parentNode);
    if (Gr && ze !== window) {
      var sn = matrix(Yr || ze), ln = sn && sn.a, un = sn && sn.d;
      sn && (Zr /= un, Kr /= ln, rn /= ln, nn /= un, en = Zr + nn, tn = Kr + rn);
    }
    return {
      top: Zr,
      left: Kr,
      bottom: en,
      right: tn,
      width: rn,
      height: nn
    };
  }
}
function isScrolledPast(ze, Qr, Wr) {
  for (var Gr = getParentAutoScrollElement(ze, !0), Yr = getRect(ze)[Qr]; Gr; ) {
    var Jr = getRect(Gr)[Wr], Zr = void 0;
    if (Zr = Yr >= Jr, !Zr) return Gr;
    if (Gr === getWindowScrollingElement()) break;
    Gr = getParentAutoScrollElement(Gr, !1);
  }
  return !1;
}
function getChild(ze, Qr, Wr, Gr) {
  for (var Yr = 0, Jr = 0, Zr = ze.children; Jr < Zr.length; ) {
    if (Zr[Jr].style.display !== "none" && Zr[Jr] !== Sortable.ghost && (Gr || Zr[Jr] !== Sortable.dragged) && closest(Zr[Jr], Wr.draggable, ze, !1)) {
      if (Yr === Qr)
        return Zr[Jr];
      Yr++;
    }
    Jr++;
  }
  return null;
}
function lastChild(ze, Qr) {
  for (var Wr = ze.lastElementChild; Wr && (Wr === Sortable.ghost || css(Wr, "display") === "none" || Qr && !matches(Wr, Qr)); )
    Wr = Wr.previousElementSibling;
  return Wr || null;
}
function index(ze, Qr) {
  var Wr = 0;
  if (!ze || !ze.parentNode)
    return -1;
  for (; ze = ze.previousElementSibling; )
    ze.nodeName.toUpperCase() !== "TEMPLATE" && ze !== Sortable.clone && (!Qr || matches(ze, Qr)) && Wr++;
  return Wr;
}
function getRelativeScrollOffset(ze) {
  var Qr = 0, Wr = 0, Gr = getWindowScrollingElement();
  if (ze)
    do {
      var Yr = matrix(ze), Jr = Yr.a, Zr = Yr.d;
      Qr += ze.scrollLeft * Jr, Wr += ze.scrollTop * Zr;
    } while (ze !== Gr && (ze = ze.parentNode));
  return [Qr, Wr];
}
function indexOfObject(ze, Qr) {
  for (var Wr in ze)
    if (ze.hasOwnProperty(Wr)) {
      for (var Gr in Qr)
        if (Qr.hasOwnProperty(Gr) && Qr[Gr] === ze[Wr][Gr]) return Number(Wr);
    }
  return -1;
}
function getParentAutoScrollElement(ze, Qr) {
  if (!ze || !ze.getBoundingClientRect) return getWindowScrollingElement();
  var Wr = ze, Gr = !1;
  do
    if (Wr.clientWidth < Wr.scrollWidth || Wr.clientHeight < Wr.scrollHeight) {
      var Yr = css(Wr);
      if (Wr.clientWidth < Wr.scrollWidth && (Yr.overflowX == "auto" || Yr.overflowX == "scroll") || Wr.clientHeight < Wr.scrollHeight && (Yr.overflowY == "auto" || Yr.overflowY == "scroll")) {
        if (!Wr.getBoundingClientRect || Wr === document.body) return getWindowScrollingElement();
        if (Gr || Qr) return Wr;
        Gr = !0;
      }
    }
  while (Wr = Wr.parentNode);
  return getWindowScrollingElement();
}
function extend(ze, Qr) {
  if (ze && Qr)
    for (var Wr in Qr)
      Qr.hasOwnProperty(Wr) && (ze[Wr] = Qr[Wr]);
  return ze;
}
function isRectEqual(ze, Qr) {
  return Math.round(ze.top) === Math.round(Qr.top) && Math.round(ze.left) === Math.round(Qr.left) && Math.round(ze.height) === Math.round(Qr.height) && Math.round(ze.width) === Math.round(Qr.width);
}
var _throttleTimeout;
function throttle(ze, Qr) {
  return function() {
    if (!_throttleTimeout) {
      var Wr = arguments, Gr = this;
      Wr.length === 1 ? ze.call(Gr, Wr[0]) : ze.apply(Gr, Wr), _throttleTimeout = setTimeout(function() {
        _throttleTimeout = void 0;
      }, Qr);
    }
  };
}
function cancelThrottle() {
  clearTimeout(_throttleTimeout), _throttleTimeout = void 0;
}
function scrollBy(ze, Qr, Wr) {
  ze.scrollLeft += Qr, ze.scrollTop += Wr;
}
function clone(ze) {
  var Qr = window.Polymer, Wr = window.jQuery || window.Zepto;
  return Qr && Qr.dom ? Qr.dom(ze).cloneNode(!0) : Wr ? Wr(ze).clone(!0)[0] : ze.cloneNode(!0);
}
function getChildContainingRectFromElement(ze, Qr, Wr) {
  var Gr = {};
  return Array.from(ze.children).forEach(function(Yr) {
    var Jr, Zr, Kr, en;
    if (!(!closest(Yr, Qr.draggable, ze, !1) || Yr.animated || Yr === Wr)) {
      var tn = getRect(Yr);
      Gr.left = Math.min((Jr = Gr.left) !== null && Jr !== void 0 ? Jr : 1 / 0, tn.left), Gr.top = Math.min((Zr = Gr.top) !== null && Zr !== void 0 ? Zr : 1 / 0, tn.top), Gr.right = Math.max((Kr = Gr.right) !== null && Kr !== void 0 ? Kr : -1 / 0, tn.right), Gr.bottom = Math.max((en = Gr.bottom) !== null && en !== void 0 ? en : -1 / 0, tn.bottom);
    }
  }), Gr.width = Gr.right - Gr.left, Gr.height = Gr.bottom - Gr.top, Gr.x = Gr.left, Gr.y = Gr.top, Gr;
}
var expando = "Sortable" + (/* @__PURE__ */ new Date()).getTime();
function AnimationStateManager() {
  var ze = [], Qr;
  return {
    captureAnimationState: function() {
      if (ze = [], !!this.options.animation) {
        var Gr = [].slice.call(this.el.children);
        Gr.forEach(function(Yr) {
          if (!(css(Yr, "display") === "none" || Yr === Sortable.ghost)) {
            ze.push({
              target: Yr,
              rect: getRect(Yr)
            });
            var Jr = _objectSpread2({}, ze[ze.length - 1].rect);
            if (Yr.thisAnimationDuration) {
              var Zr = matrix(Yr, !0);
              Zr && (Jr.top -= Zr.f, Jr.left -= Zr.e);
            }
            Yr.fromRect = Jr;
          }
        });
      }
    },
    addAnimationState: function(Gr) {
      ze.push(Gr);
    },
    removeAnimationState: function(Gr) {
      ze.splice(indexOfObject(ze, {
        target: Gr
      }), 1);
    },
    animateAll: function(Gr) {
      var Yr = this;
      if (!this.options.animation) {
        clearTimeout(Qr), typeof Gr == "function" && Gr();
        return;
      }
      var Jr = !1, Zr = 0;
      ze.forEach(function(Kr) {
        var en = 0, tn = Kr.target, nn = tn.fromRect, rn = getRect(tn), an = tn.prevFromRect, sn = tn.prevToRect, ln = Kr.rect, un = matrix(tn, !0);
        un && (rn.top -= un.f, rn.left -= un.e), tn.toRect = rn, tn.thisAnimationDuration && isRectEqual(an, rn) && !isRectEqual(nn, rn) && // Make sure animatingRect is on line between toRect & fromRect
        (ln.top - rn.top) / (ln.left - rn.left) === (nn.top - rn.top) / (nn.left - rn.left) && (en = calculateRealTime(ln, an, sn, Yr.options)), isRectEqual(rn, nn) || (tn.prevFromRect = nn, tn.prevToRect = rn, en || (en = Yr.options.animation), Yr.animate(tn, ln, rn, en)), en && (Jr = !0, Zr = Math.max(Zr, en), clearTimeout(tn.animationResetTimer), tn.animationResetTimer = setTimeout(function() {
          tn.animationTime = 0, tn.prevFromRect = null, tn.fromRect = null, tn.prevToRect = null, tn.thisAnimationDuration = null;
        }, en), tn.thisAnimationDuration = en);
      }), clearTimeout(Qr), Jr ? Qr = setTimeout(function() {
        typeof Gr == "function" && Gr();
      }, Zr) : typeof Gr == "function" && Gr(), ze = [];
    },
    animate: function(Gr, Yr, Jr, Zr) {
      if (Zr) {
        css(Gr, "transition", ""), css(Gr, "transform", "");
        var Kr = matrix(this.el), en = Kr && Kr.a, tn = Kr && Kr.d, nn = (Yr.left - Jr.left) / (en || 1), rn = (Yr.top - Jr.top) / (tn || 1);
        Gr.animatingX = !!nn, Gr.animatingY = !!rn, css(Gr, "transform", "translate3d(" + nn + "px," + rn + "px,0)"), this.forRepaintDummy = repaint(Gr), css(Gr, "transition", "transform " + Zr + "ms" + (this.options.easing ? " " + this.options.easing : "")), css(Gr, "transform", "translate3d(0,0,0)"), typeof Gr.animated == "number" && clearTimeout(Gr.animated), Gr.animated = setTimeout(function() {
          css(Gr, "transition", ""), css(Gr, "transform", ""), Gr.animated = !1, Gr.animatingX = !1, Gr.animatingY = !1;
        }, Zr);
      }
    }
  };
}
function repaint(ze) {
  return ze.offsetWidth;
}
function calculateRealTime(ze, Qr, Wr, Gr) {
  return Math.sqrt(Math.pow(Qr.top - ze.top, 2) + Math.pow(Qr.left - ze.left, 2)) / Math.sqrt(Math.pow(Qr.top - Wr.top, 2) + Math.pow(Qr.left - Wr.left, 2)) * Gr.animation;
}
var plugins = [], defaults = {
  initializeByDefault: !0
}, PluginManager = {
  mount: function ze(Qr) {
    for (var Wr in defaults)
      defaults.hasOwnProperty(Wr) && !(Wr in Qr) && (Qr[Wr] = defaults[Wr]);
    plugins.forEach(function(Gr) {
      if (Gr.pluginName === Qr.pluginName)
        throw "Sortable: Cannot mount plugin ".concat(Qr.pluginName, " more than once");
    }), plugins.push(Qr);
  },
  pluginEvent: function ze(Qr, Wr, Gr) {
    var Yr = this;
    this.eventCanceled = !1, Gr.cancel = function() {
      Yr.eventCanceled = !0;
    };
    var Jr = Qr + "Global";
    plugins.forEach(function(Zr) {
      Wr[Zr.pluginName] && (Wr[Zr.pluginName][Jr] && Wr[Zr.pluginName][Jr](_objectSpread2({
        sortable: Wr
      }, Gr)), Wr.options[Zr.pluginName] && Wr[Zr.pluginName][Qr] && Wr[Zr.pluginName][Qr](_objectSpread2({
        sortable: Wr
      }, Gr)));
    });
  },
  initializePlugins: function ze(Qr, Wr, Gr, Yr) {
    plugins.forEach(function(Kr) {
      var en = Kr.pluginName;
      if (!(!Qr.options[en] && !Kr.initializeByDefault)) {
        var tn = new Kr(Qr, Wr, Qr.options);
        tn.sortable = Qr, tn.options = Qr.options, Qr[en] = tn, _extends(Gr, tn.defaults);
      }
    });
    for (var Jr in Qr.options)
      if (Qr.options.hasOwnProperty(Jr)) {
        var Zr = this.modifyOption(Qr, Jr, Qr.options[Jr]);
        typeof Zr < "u" && (Qr.options[Jr] = Zr);
      }
  },
  getEventProperties: function ze(Qr, Wr) {
    var Gr = {};
    return plugins.forEach(function(Yr) {
      typeof Yr.eventProperties == "function" && _extends(Gr, Yr.eventProperties.call(Wr[Yr.pluginName], Qr));
    }), Gr;
  },
  modifyOption: function ze(Qr, Wr, Gr) {
    var Yr;
    return plugins.forEach(function(Jr) {
      Qr[Jr.pluginName] && Jr.optionListeners && typeof Jr.optionListeners[Wr] == "function" && (Yr = Jr.optionListeners[Wr].call(Qr[Jr.pluginName], Gr));
    }), Yr;
  }
};
function dispatchEvent(ze) {
  var Qr = ze.sortable, Wr = ze.rootEl, Gr = ze.name, Yr = ze.targetEl, Jr = ze.cloneEl, Zr = ze.toEl, Kr = ze.fromEl, en = ze.oldIndex, tn = ze.newIndex, nn = ze.oldDraggableIndex, rn = ze.newDraggableIndex, an = ze.originalEvent, sn = ze.putSortable, ln = ze.extraEventProperties;
  if (Qr = Qr || Wr && Wr[expando], !!Qr) {
    var un, bn = Qr.options, Cn = "on" + Gr.charAt(0).toUpperCase() + Gr.substr(1);
    window.CustomEvent && !IE11OrLess && !Edge ? un = new CustomEvent(Gr, {
      bubbles: !0,
      cancelable: !0
    }) : (un = document.createEvent("Event"), un.initEvent(Gr, !0, !0)), un.to = Zr || Wr, un.from = Kr || Wr, un.item = Yr || Wr, un.clone = Jr, un.oldIndex = en, un.newIndex = tn, un.oldDraggableIndex = nn, un.newDraggableIndex = rn, un.originalEvent = an, un.pullMode = sn ? sn.lastPutMode : void 0;
    var hn = _objectSpread2(_objectSpread2({}, ln), PluginManager.getEventProperties(Gr, Qr));
    for (var gn in hn)
      un[gn] = hn[gn];
    Wr && Wr.dispatchEvent(un), bn[Cn] && bn[Cn].call(Qr, un);
  }
}
var _excluded = ["evt"], pluginEvent = function ze(Qr, Wr) {
  var Gr = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : {}, Yr = Gr.evt, Jr = _objectWithoutProperties(Gr, _excluded);
  PluginManager.pluginEvent.bind(Sortable)(Qr, Wr, _objectSpread2({
    dragEl,
    parentEl,
    ghostEl,
    rootEl,
    nextEl,
    lastDownEl,
    cloneEl,
    cloneHidden,
    dragStarted: moved,
    putSortable,
    activeSortable: Sortable.active,
    originalEvent: Yr,
    oldIndex,
    oldDraggableIndex,
    newIndex,
    newDraggableIndex,
    hideGhostForTarget: _hideGhostForTarget,
    unhideGhostForTarget: _unhideGhostForTarget,
    cloneNowHidden: function() {
      cloneHidden = !0;
    },
    cloneNowShown: function() {
      cloneHidden = !1;
    },
    dispatchSortableEvent: function(Kr) {
      _dispatchEvent({
        sortable: Wr,
        name: Kr,
        originalEvent: Yr
      });
    }
  }, Jr));
};
function _dispatchEvent(ze) {
  dispatchEvent(_objectSpread2({
    putSortable,
    cloneEl,
    targetEl: dragEl,
    rootEl,
    oldIndex,
    oldDraggableIndex,
    newIndex,
    newDraggableIndex
  }, ze));
}
var dragEl, parentEl, ghostEl, rootEl, nextEl, lastDownEl, cloneEl, cloneHidden, oldIndex, newIndex, oldDraggableIndex, newDraggableIndex, activeGroup, putSortable, awaitingDragStarted = !1, ignoreNextClick = !1, sortables = [], tapEvt, touchEvt, lastDx, lastDy, tapDistanceLeft, tapDistanceTop, moved, lastTarget, lastDirection, pastFirstInvertThresh = !1, isCircumstantialInvert = !1, targetMoveDistance, ghostRelativeParent, ghostRelativeParentInitialScroll = [], _silent = !1, savedInputChecked = [], documentExists = typeof document < "u", PositionGhostAbsolutely = IOS, CSSFloatProperty = Edge || IE11OrLess ? "cssFloat" : "float", supportDraggable = documentExists && !ChromeForAndroid && !IOS && "draggable" in document.createElement("div"), supportCssPointerEvents = function() {
  if (documentExists) {
    if (IE11OrLess)
      return !1;
    var ze = document.createElement("x");
    return ze.style.cssText = "pointer-events:auto", ze.style.pointerEvents === "auto";
  }
}(), _detectDirection = function ze(Qr, Wr) {
  var Gr = css(Qr), Yr = parseInt(Gr.width) - parseInt(Gr.paddingLeft) - parseInt(Gr.paddingRight) - parseInt(Gr.borderLeftWidth) - parseInt(Gr.borderRightWidth), Jr = getChild(Qr, 0, Wr), Zr = getChild(Qr, 1, Wr), Kr = Jr && css(Jr), en = Zr && css(Zr), tn = Kr && parseInt(Kr.marginLeft) + parseInt(Kr.marginRight) + getRect(Jr).width, nn = en && parseInt(en.marginLeft) + parseInt(en.marginRight) + getRect(Zr).width;
  if (Gr.display === "flex")
    return Gr.flexDirection === "column" || Gr.flexDirection === "column-reverse" ? "vertical" : "horizontal";
  if (Gr.display === "grid")
    return Gr.gridTemplateColumns.split(" ").length <= 1 ? "vertical" : "horizontal";
  if (Jr && Kr.float && Kr.float !== "none") {
    var rn = Kr.float === "left" ? "left" : "right";
    return Zr && (en.clear === "both" || en.clear === rn) ? "vertical" : "horizontal";
  }
  return Jr && (Kr.display === "block" || Kr.display === "flex" || Kr.display === "table" || Kr.display === "grid" || tn >= Yr && Gr[CSSFloatProperty] === "none" || Zr && Gr[CSSFloatProperty] === "none" && tn + nn > Yr) ? "vertical" : "horizontal";
}, _dragElInRowColumn = function ze(Qr, Wr, Gr) {
  var Yr = Gr ? Qr.left : Qr.top, Jr = Gr ? Qr.right : Qr.bottom, Zr = Gr ? Qr.width : Qr.height, Kr = Gr ? Wr.left : Wr.top, en = Gr ? Wr.right : Wr.bottom, tn = Gr ? Wr.width : Wr.height;
  return Yr === Kr || Jr === en || Yr + Zr / 2 === Kr + tn / 2;
}, _detectNearestEmptySortable = function ze(Qr, Wr) {
  var Gr;
  return sortables.some(function(Yr) {
    var Jr = Yr[expando].options.emptyInsertThreshold;
    if (!(!Jr || lastChild(Yr))) {
      var Zr = getRect(Yr), Kr = Qr >= Zr.left - Jr && Qr <= Zr.right + Jr, en = Wr >= Zr.top - Jr && Wr <= Zr.bottom + Jr;
      if (Kr && en)
        return Gr = Yr;
    }
  }), Gr;
}, _prepareGroup = function ze(Qr) {
  function Wr(Jr, Zr) {
    return function(Kr, en, tn, nn) {
      var rn = Kr.options.group.name && en.options.group.name && Kr.options.group.name === en.options.group.name;
      if (Jr == null && (Zr || rn))
        return !0;
      if (Jr == null || Jr === !1)
        return !1;
      if (Zr && Jr === "clone")
        return Jr;
      if (typeof Jr == "function")
        return Wr(Jr(Kr, en, tn, nn), Zr)(Kr, en, tn, nn);
      var an = (Zr ? Kr : en).options.group.name;
      return Jr === !0 || typeof Jr == "string" && Jr === an || Jr.join && Jr.indexOf(an) > -1;
    };
  }
  var Gr = {}, Yr = Qr.group;
  (!Yr || _typeof(Yr) != "object") && (Yr = {
    name: Yr
  }), Gr.name = Yr.name, Gr.checkPull = Wr(Yr.pull, !0), Gr.checkPut = Wr(Yr.put), Gr.revertClone = Yr.revertClone, Qr.group = Gr;
}, _hideGhostForTarget = function ze() {
  !supportCssPointerEvents && ghostEl && css(ghostEl, "display", "none");
}, _unhideGhostForTarget = function ze() {
  !supportCssPointerEvents && ghostEl && css(ghostEl, "display", "");
};
documentExists && !ChromeForAndroid && document.addEventListener("click", function(ze) {
  if (ignoreNextClick)
    return ze.preventDefault(), ze.stopPropagation && ze.stopPropagation(), ze.stopImmediatePropagation && ze.stopImmediatePropagation(), ignoreNextClick = !1, !1;
}, !0);
var nearestEmptyInsertDetectEvent = function ze(Qr) {
  if (dragEl) {
    Qr = Qr.touches ? Qr.touches[0] : Qr;
    var Wr = _detectNearestEmptySortable(Qr.clientX, Qr.clientY);
    if (Wr) {
      var Gr = {};
      for (var Yr in Qr)
        Qr.hasOwnProperty(Yr) && (Gr[Yr] = Qr[Yr]);
      Gr.target = Gr.rootEl = Wr, Gr.preventDefault = void 0, Gr.stopPropagation = void 0, Wr[expando]._onDragOver(Gr);
    }
  }
}, _checkOutsideTargetEl = function ze(Qr) {
  dragEl && dragEl.parentNode[expando]._isOutsideThisEl(Qr.target);
};
function Sortable(ze, Qr) {
  if (!(ze && ze.nodeType && ze.nodeType === 1))
    throw "Sortable: `el` must be an HTMLElement, not ".concat({}.toString.call(ze));
  this.el = ze, this.options = Qr = _extends({}, Qr), ze[expando] = this;
  var Wr = {
    group: null,
    sort: !0,
    disabled: !1,
    store: null,
    handle: null,
    draggable: /^[uo]l$/i.test(ze.nodeName) ? ">li" : ">*",
    swapThreshold: 1,
    // percentage; 0 <= x <= 1
    invertSwap: !1,
    // invert always
    invertedSwapThreshold: null,
    // will be set to same as swapThreshold if default
    removeCloneOnHide: !0,
    direction: function() {
      return _detectDirection(ze, this.options);
    },
    ghostClass: "sortable-ghost",
    chosenClass: "sortable-chosen",
    dragClass: "sortable-drag",
    ignore: "a, img",
    filter: null,
    preventOnFilter: !0,
    animation: 0,
    easing: null,
    setData: function(Zr, Kr) {
      Zr.setData("Text", Kr.textContent);
    },
    dropBubble: !1,
    dragoverBubble: !1,
    dataIdAttr: "data-id",
    delay: 0,
    delayOnTouchOnly: !1,
    touchStartThreshold: (Number.parseInt ? Number : window).parseInt(window.devicePixelRatio, 10) || 1,
    forceFallback: !1,
    fallbackClass: "sortable-fallback",
    fallbackOnBody: !1,
    fallbackTolerance: 0,
    fallbackOffset: {
      x: 0,
      y: 0
    },
    // Disabled on Safari: #1571; Enabled on Safari IOS: #2244
    supportPointer: Sortable.supportPointer !== !1 && "PointerEvent" in window && (!Safari || IOS),
    emptyInsertThreshold: 5
  };
  PluginManager.initializePlugins(this, ze, Wr);
  for (var Gr in Wr)
    !(Gr in Qr) && (Qr[Gr] = Wr[Gr]);
  _prepareGroup(Qr);
  for (var Yr in this)
    Yr.charAt(0) === "_" && typeof this[Yr] == "function" && (this[Yr] = this[Yr].bind(this));
  this.nativeDraggable = Qr.forceFallback ? !1 : supportDraggable, this.nativeDraggable && (this.options.touchStartThreshold = 1), Qr.supportPointer ? on(ze, "pointerdown", this._onTapStart) : (on(ze, "mousedown", this._onTapStart), on(ze, "touchstart", this._onTapStart)), this.nativeDraggable && (on(ze, "dragover", this), on(ze, "dragenter", this)), sortables.push(this.el), Qr.store && Qr.store.get && this.sort(Qr.store.get(this) || []), _extends(this, AnimationStateManager());
}
Sortable.prototype = /** @lends Sortable.prototype */
{
  constructor: Sortable,
  _isOutsideThisEl: function ze(Qr) {
    !this.el.contains(Qr) && Qr !== this.el && (lastTarget = null);
  },
  _getDirection: function ze(Qr, Wr) {
    return typeof this.options.direction == "function" ? this.options.direction.call(this, Qr, Wr, dragEl) : this.options.direction;
  },
  _onTapStart: function ze(Qr) {
    if (Qr.cancelable) {
      var Wr = this, Gr = this.el, Yr = this.options, Jr = Yr.preventOnFilter, Zr = Qr.type, Kr = Qr.touches && Qr.touches[0] || Qr.pointerType && Qr.pointerType === "touch" && Qr, en = (Kr || Qr).target, tn = Qr.target.shadowRoot && (Qr.path && Qr.path[0] || Qr.composedPath && Qr.composedPath()[0]) || en, nn = Yr.filter;
      if (_saveInputCheckedState(Gr), !dragEl && !(/mousedown|pointerdown/.test(Zr) && Qr.button !== 0 || Yr.disabled) && !tn.isContentEditable && !(!this.nativeDraggable && Safari && en && en.tagName.toUpperCase() === "SELECT") && (en = closest(en, Yr.draggable, Gr, !1), !(en && en.animated) && lastDownEl !== en)) {
        if (oldIndex = index(en), oldDraggableIndex = index(en, Yr.draggable), typeof nn == "function") {
          if (nn.call(this, Qr, en, this)) {
            _dispatchEvent({
              sortable: Wr,
              rootEl: tn,
              name: "filter",
              targetEl: en,
              toEl: Gr,
              fromEl: Gr
            }), pluginEvent("filter", Wr, {
              evt: Qr
            }), Jr && Qr.preventDefault();
            return;
          }
        } else if (nn && (nn = nn.split(",").some(function(rn) {
          if (rn = closest(tn, rn.trim(), Gr, !1), rn)
            return _dispatchEvent({
              sortable: Wr,
              rootEl: rn,
              name: "filter",
              targetEl: en,
              fromEl: Gr,
              toEl: Gr
            }), pluginEvent("filter", Wr, {
              evt: Qr
            }), !0;
        }), nn)) {
          Jr && Qr.preventDefault();
          return;
        }
        Yr.handle && !closest(tn, Yr.handle, Gr, !1) || this._prepareDragStart(Qr, Kr, en);
      }
    }
  },
  _prepareDragStart: function ze(Qr, Wr, Gr) {
    var Yr = this, Jr = Yr.el, Zr = Yr.options, Kr = Jr.ownerDocument, en;
    if (Gr && !dragEl && Gr.parentNode === Jr) {
      var tn = getRect(Gr);
      if (rootEl = Jr, dragEl = Gr, parentEl = dragEl.parentNode, nextEl = dragEl.nextSibling, lastDownEl = Gr, activeGroup = Zr.group, Sortable.dragged = dragEl, tapEvt = {
        target: dragEl,
        clientX: (Wr || Qr).clientX,
        clientY: (Wr || Qr).clientY
      }, tapDistanceLeft = tapEvt.clientX - tn.left, tapDistanceTop = tapEvt.clientY - tn.top, this._lastX = (Wr || Qr).clientX, this._lastY = (Wr || Qr).clientY, dragEl.style["will-change"] = "all", en = function() {
        if (pluginEvent("delayEnded", Yr, {
          evt: Qr
        }), Sortable.eventCanceled) {
          Yr._onDrop();
          return;
        }
        Yr._disableDelayedDragEvents(), !FireFox && Yr.nativeDraggable && (dragEl.draggable = !0), Yr._triggerDragStart(Qr, Wr), _dispatchEvent({
          sortable: Yr,
          name: "choose",
          originalEvent: Qr
        }), toggleClass(dragEl, Zr.chosenClass, !0);
      }, Zr.ignore.split(",").forEach(function(nn) {
        find(dragEl, nn.trim(), _disableDraggable);
      }), on(Kr, "dragover", nearestEmptyInsertDetectEvent), on(Kr, "mousemove", nearestEmptyInsertDetectEvent), on(Kr, "touchmove", nearestEmptyInsertDetectEvent), Zr.supportPointer ? (on(Kr, "pointerup", Yr._onDrop), !this.nativeDraggable && on(Kr, "pointercancel", Yr._onDrop)) : (on(Kr, "mouseup", Yr._onDrop), on(Kr, "touchend", Yr._onDrop), on(Kr, "touchcancel", Yr._onDrop)), FireFox && this.nativeDraggable && (this.options.touchStartThreshold = 4, dragEl.draggable = !0), pluginEvent("delayStart", this, {
        evt: Qr
      }), Zr.delay && (!Zr.delayOnTouchOnly || Wr) && (!this.nativeDraggable || !(Edge || IE11OrLess))) {
        if (Sortable.eventCanceled) {
          this._onDrop();
          return;
        }
        Zr.supportPointer ? (on(Kr, "pointerup", Yr._disableDelayedDrag), on(Kr, "pointercancel", Yr._disableDelayedDrag)) : (on(Kr, "mouseup", Yr._disableDelayedDrag), on(Kr, "touchend", Yr._disableDelayedDrag), on(Kr, "touchcancel", Yr._disableDelayedDrag)), on(Kr, "mousemove", Yr._delayedDragTouchMoveHandler), on(Kr, "touchmove", Yr._delayedDragTouchMoveHandler), Zr.supportPointer && on(Kr, "pointermove", Yr._delayedDragTouchMoveHandler), Yr._dragStartTimer = setTimeout(en, Zr.delay);
      } else
        en();
    }
  },
  _delayedDragTouchMoveHandler: function ze(Qr) {
    var Wr = Qr.touches ? Qr.touches[0] : Qr;
    Math.max(Math.abs(Wr.clientX - this._lastX), Math.abs(Wr.clientY - this._lastY)) >= Math.floor(this.options.touchStartThreshold / (this.nativeDraggable && window.devicePixelRatio || 1)) && this._disableDelayedDrag();
  },
  _disableDelayedDrag: function ze() {
    dragEl && _disableDraggable(dragEl), clearTimeout(this._dragStartTimer), this._disableDelayedDragEvents();
  },
  _disableDelayedDragEvents: function ze() {
    var Qr = this.el.ownerDocument;
    off(Qr, "mouseup", this._disableDelayedDrag), off(Qr, "touchend", this._disableDelayedDrag), off(Qr, "touchcancel", this._disableDelayedDrag), off(Qr, "pointerup", this._disableDelayedDrag), off(Qr, "pointercancel", this._disableDelayedDrag), off(Qr, "mousemove", this._delayedDragTouchMoveHandler), off(Qr, "touchmove", this._delayedDragTouchMoveHandler), off(Qr, "pointermove", this._delayedDragTouchMoveHandler);
  },
  _triggerDragStart: function ze(Qr, Wr) {
    Wr = Wr || Qr.pointerType == "touch" && Qr, !this.nativeDraggable || Wr ? this.options.supportPointer ? on(document, "pointermove", this._onTouchMove) : Wr ? on(document, "touchmove", this._onTouchMove) : on(document, "mousemove", this._onTouchMove) : (on(dragEl, "dragend", this), on(rootEl, "dragstart", this._onDragStart));
    try {
      document.selection ? _nextTick(function() {
        document.selection.empty();
      }) : window.getSelection().removeAllRanges();
    } catch {
    }
  },
  _dragStarted: function ze(Qr, Wr) {
    if (awaitingDragStarted = !1, rootEl && dragEl) {
      pluginEvent("dragStarted", this, {
        evt: Wr
      }), this.nativeDraggable && on(document, "dragover", _checkOutsideTargetEl);
      var Gr = this.options;
      !Qr && toggleClass(dragEl, Gr.dragClass, !1), toggleClass(dragEl, Gr.ghostClass, !0), Sortable.active = this, Qr && this._appendGhost(), _dispatchEvent({
        sortable: this,
        name: "start",
        originalEvent: Wr
      });
    } else
      this._nulling();
  },
  _emulateDragOver: function ze() {
    if (touchEvt) {
      this._lastX = touchEvt.clientX, this._lastY = touchEvt.clientY, _hideGhostForTarget();
      for (var Qr = document.elementFromPoint(touchEvt.clientX, touchEvt.clientY), Wr = Qr; Qr && Qr.shadowRoot && (Qr = Qr.shadowRoot.elementFromPoint(touchEvt.clientX, touchEvt.clientY), Qr !== Wr); )
        Wr = Qr;
      if (dragEl.parentNode[expando]._isOutsideThisEl(Qr), Wr)
        do {
          if (Wr[expando]) {
            var Gr = void 0;
            if (Gr = Wr[expando]._onDragOver({
              clientX: touchEvt.clientX,
              clientY: touchEvt.clientY,
              target: Qr,
              rootEl: Wr
            }), Gr && !this.options.dragoverBubble)
              break;
          }
          Qr = Wr;
        } while (Wr = getParentOrHost(Wr));
      _unhideGhostForTarget();
    }
  },
  _onTouchMove: function ze(Qr) {
    if (tapEvt) {
      var Wr = this.options, Gr = Wr.fallbackTolerance, Yr = Wr.fallbackOffset, Jr = Qr.touches ? Qr.touches[0] : Qr, Zr = ghostEl && matrix(ghostEl, !0), Kr = ghostEl && Zr && Zr.a, en = ghostEl && Zr && Zr.d, tn = PositionGhostAbsolutely && ghostRelativeParent && getRelativeScrollOffset(ghostRelativeParent), nn = (Jr.clientX - tapEvt.clientX + Yr.x) / (Kr || 1) + (tn ? tn[0] - ghostRelativeParentInitialScroll[0] : 0) / (Kr || 1), rn = (Jr.clientY - tapEvt.clientY + Yr.y) / (en || 1) + (tn ? tn[1] - ghostRelativeParentInitialScroll[1] : 0) / (en || 1);
      if (!Sortable.active && !awaitingDragStarted) {
        if (Gr && Math.max(Math.abs(Jr.clientX - this._lastX), Math.abs(Jr.clientY - this._lastY)) < Gr)
          return;
        this._onDragStart(Qr, !0);
      }
      if (ghostEl) {
        Zr ? (Zr.e += nn - (lastDx || 0), Zr.f += rn - (lastDy || 0)) : Zr = {
          a: 1,
          b: 0,
          c: 0,
          d: 1,
          e: nn,
          f: rn
        };
        var an = "matrix(".concat(Zr.a, ",").concat(Zr.b, ",").concat(Zr.c, ",").concat(Zr.d, ",").concat(Zr.e, ",").concat(Zr.f, ")");
        css(ghostEl, "webkitTransform", an), css(ghostEl, "mozTransform", an), css(ghostEl, "msTransform", an), css(ghostEl, "transform", an), lastDx = nn, lastDy = rn, touchEvt = Jr;
      }
      Qr.cancelable && Qr.preventDefault();
    }
  },
  _appendGhost: function ze() {
    if (!ghostEl) {
      var Qr = this.options.fallbackOnBody ? document.body : rootEl, Wr = getRect(dragEl, !0, PositionGhostAbsolutely, !0, Qr), Gr = this.options;
      if (PositionGhostAbsolutely) {
        for (ghostRelativeParent = Qr; css(ghostRelativeParent, "position") === "static" && css(ghostRelativeParent, "transform") === "none" && ghostRelativeParent !== document; )
          ghostRelativeParent = ghostRelativeParent.parentNode;
        ghostRelativeParent !== document.body && ghostRelativeParent !== document.documentElement ? (ghostRelativeParent === document && (ghostRelativeParent = getWindowScrollingElement()), Wr.top += ghostRelativeParent.scrollTop, Wr.left += ghostRelativeParent.scrollLeft) : ghostRelativeParent = getWindowScrollingElement(), ghostRelativeParentInitialScroll = getRelativeScrollOffset(ghostRelativeParent);
      }
      ghostEl = dragEl.cloneNode(!0), toggleClass(ghostEl, Gr.ghostClass, !1), toggleClass(ghostEl, Gr.fallbackClass, !0), toggleClass(ghostEl, Gr.dragClass, !0), css(ghostEl, "transition", ""), css(ghostEl, "transform", ""), css(ghostEl, "box-sizing", "border-box"), css(ghostEl, "margin", 0), css(ghostEl, "top", Wr.top), css(ghostEl, "left", Wr.left), css(ghostEl, "width", Wr.width), css(ghostEl, "height", Wr.height), css(ghostEl, "opacity", "0.8"), css(ghostEl, "position", PositionGhostAbsolutely ? "absolute" : "fixed"), css(ghostEl, "zIndex", "100000"), css(ghostEl, "pointerEvents", "none"), Sortable.ghost = ghostEl, Qr.appendChild(ghostEl), css(ghostEl, "transform-origin", tapDistanceLeft / parseInt(ghostEl.style.width) * 100 + "% " + tapDistanceTop / parseInt(ghostEl.style.height) * 100 + "%");
    }
  },
  _onDragStart: function ze(Qr, Wr) {
    var Gr = this, Yr = Qr.dataTransfer, Jr = Gr.options;
    if (pluginEvent("dragStart", this, {
      evt: Qr
    }), Sortable.eventCanceled) {
      this._onDrop();
      return;
    }
    pluginEvent("setupClone", this), Sortable.eventCanceled || (cloneEl = clone(dragEl), cloneEl.removeAttribute("id"), cloneEl.draggable = !1, cloneEl.style["will-change"] = "", this._hideClone(), toggleClass(cloneEl, this.options.chosenClass, !1), Sortable.clone = cloneEl), Gr.cloneId = _nextTick(function() {
      pluginEvent("clone", Gr), !Sortable.eventCanceled && (Gr.options.removeCloneOnHide || rootEl.insertBefore(cloneEl, dragEl), Gr._hideClone(), _dispatchEvent({
        sortable: Gr,
        name: "clone"
      }));
    }), !Wr && toggleClass(dragEl, Jr.dragClass, !0), Wr ? (ignoreNextClick = !0, Gr._loopId = setInterval(Gr._emulateDragOver, 50)) : (off(document, "mouseup", Gr._onDrop), off(document, "touchend", Gr._onDrop), off(document, "touchcancel", Gr._onDrop), Yr && (Yr.effectAllowed = "move", Jr.setData && Jr.setData.call(Gr, Yr, dragEl)), on(document, "drop", Gr), css(dragEl, "transform", "translateZ(0)")), awaitingDragStarted = !0, Gr._dragStartId = _nextTick(Gr._dragStarted.bind(Gr, Wr, Qr)), on(document, "selectstart", Gr), moved = !0, window.getSelection().removeAllRanges(), Safari && css(document.body, "user-select", "none");
  },
  // Returns true - if no further action is needed (either inserted or another condition)
  _onDragOver: function ze(Qr) {
    var Wr = this.el, Gr = Qr.target, Yr, Jr, Zr, Kr = this.options, en = Kr.group, tn = Sortable.active, nn = activeGroup === en, rn = Kr.sort, an = putSortable || tn, sn, ln = this, un = !1;
    if (_silent) return;
    function bn(wn, Xn) {
      pluginEvent(wn, ln, _objectSpread2({
        evt: Qr,
        isOwner: nn,
        axis: sn ? "vertical" : "horizontal",
        revert: Zr,
        dragRect: Yr,
        targetRect: Jr,
        canSort: rn,
        fromSortable: an,
        target: Gr,
        completed: hn,
        onMove: function(Pn, Mn) {
          return _onMove(rootEl, Wr, dragEl, Yr, Pn, getRect(Pn), Qr, Mn);
        },
        changed: gn
      }, Xn));
    }
    function Cn() {
      bn("dragOverAnimationCapture"), ln.captureAnimationState(), ln !== an && an.captureAnimationState();
    }
    function hn(wn) {
      return bn("dragOverCompleted", {
        insertion: wn
      }), wn && (nn ? tn._hideClone() : tn._showClone(ln), ln !== an && (toggleClass(dragEl, putSortable ? putSortable.options.ghostClass : tn.options.ghostClass, !1), toggleClass(dragEl, Kr.ghostClass, !0)), putSortable !== ln && ln !== Sortable.active ? putSortable = ln : ln === Sortable.active && putSortable && (putSortable = null), an === ln && (ln._ignoreWhileAnimating = Gr), ln.animateAll(function() {
        bn("dragOverAnimationComplete"), ln._ignoreWhileAnimating = null;
      }), ln !== an && (an.animateAll(), an._ignoreWhileAnimating = null)), (Gr === dragEl && !dragEl.animated || Gr === Wr && !Gr.animated) && (lastTarget = null), !Kr.dragoverBubble && !Qr.rootEl && Gr !== document && (dragEl.parentNode[expando]._isOutsideThisEl(Qr.target), !wn && nearestEmptyInsertDetectEvent(Qr)), !Kr.dragoverBubble && Qr.stopPropagation && Qr.stopPropagation(), un = !0;
    }
    function gn() {
      newIndex = index(dragEl), newDraggableIndex = index(dragEl, Kr.draggable), _dispatchEvent({
        sortable: ln,
        name: "change",
        toEl: Wr,
        newIndex,
        newDraggableIndex,
        originalEvent: Qr
      });
    }
    if (Qr.preventDefault !== void 0 && Qr.cancelable && Qr.preventDefault(), Gr = closest(Gr, Kr.draggable, Wr, !0), bn("dragOver"), Sortable.eventCanceled) return un;
    if (dragEl.contains(Qr.target) || Gr.animated && Gr.animatingX && Gr.animatingY || ln._ignoreWhileAnimating === Gr)
      return hn(!1);
    if (ignoreNextClick = !1, tn && !Kr.disabled && (nn ? rn || (Zr = parentEl !== rootEl) : putSortable === this || (this.lastPutMode = activeGroup.checkPull(this, tn, dragEl, Qr)) && en.checkPut(this, tn, dragEl, Qr))) {
      if (sn = this._getDirection(Qr, Gr) === "vertical", Yr = getRect(dragEl), bn("dragOverValid"), Sortable.eventCanceled) return un;
      if (Zr)
        return parentEl = rootEl, Cn(), this._hideClone(), bn("revert"), Sortable.eventCanceled || (nextEl ? rootEl.insertBefore(dragEl, nextEl) : rootEl.appendChild(dragEl)), hn(!0);
      var dn = lastChild(Wr, Kr.draggable);
      if (!dn || _ghostIsLast(Qr, sn, this) && !dn.animated) {
        if (dn === dragEl)
          return hn(!1);
        if (dn && Wr === Qr.target && (Gr = dn), Gr && (Jr = getRect(Gr)), _onMove(rootEl, Wr, dragEl, Yr, Gr, Jr, Qr, !!Gr) !== !1)
          return Cn(), dn && dn.nextSibling ? Wr.insertBefore(dragEl, dn.nextSibling) : Wr.appendChild(dragEl), parentEl = Wr, gn(), hn(!0);
      } else if (dn && _ghostIsFirst(Qr, sn, this)) {
        var yn = getChild(Wr, 0, Kr, !0);
        if (yn === dragEl)
          return hn(!1);
        if (Gr = yn, Jr = getRect(Gr), _onMove(rootEl, Wr, dragEl, Yr, Gr, Jr, Qr, !1) !== !1)
          return Cn(), Wr.insertBefore(dragEl, yn), parentEl = Wr, gn(), hn(!0);
      } else if (Gr.parentNode === Wr) {
        Jr = getRect(Gr);
        var In = 0, pn, vn = dragEl.parentNode !== Wr, mn = !_dragElInRowColumn(dragEl.animated && dragEl.toRect || Yr, Gr.animated && Gr.toRect || Jr, sn), En = sn ? "top" : "left", cn = isScrolledPast(Gr, "top", "top") || isScrolledPast(dragEl, "top", "top"), Sn = cn ? cn.scrollTop : void 0;
        lastTarget !== Gr && (pn = Jr[En], pastFirstInvertThresh = !1, isCircumstantialInvert = !mn && Kr.invertSwap || vn), In = _getSwapDirection(Qr, Gr, Jr, sn, mn ? 1 : Kr.swapThreshold, Kr.invertedSwapThreshold == null ? Kr.swapThreshold : Kr.invertedSwapThreshold, isCircumstantialInvert, lastTarget === Gr);
        var xn;
        if (In !== 0) {
          var Dn = index(dragEl);
          do
            Dn -= In, xn = parentEl.children[Dn];
          while (xn && (css(xn, "display") === "none" || xn === ghostEl));
        }
        if (In === 0 || xn === Gr)
          return hn(!1);
        lastTarget = Gr, lastDirection = In;
        var fn = Gr.nextElementSibling, Tn = !1;
        Tn = In === 1;
        var Rn = _onMove(rootEl, Wr, dragEl, Yr, Gr, Jr, Qr, Tn);
        if (Rn !== !1)
          return (Rn === 1 || Rn === -1) && (Tn = Rn === 1), _silent = !0, setTimeout(_unsilent, 30), Cn(), Tn && !fn ? Wr.appendChild(dragEl) : Gr.parentNode.insertBefore(dragEl, Tn ? fn : Gr), cn && scrollBy(cn, 0, Sn - cn.scrollTop), parentEl = dragEl.parentNode, pn !== void 0 && !isCircumstantialInvert && (targetMoveDistance = Math.abs(pn - getRect(Gr)[En])), gn(), hn(!0);
      }
      if (Wr.contains(dragEl))
        return hn(!1);
    }
    return !1;
  },
  _ignoreWhileAnimating: null,
  _offMoveEvents: function ze() {
    off(document, "mousemove", this._onTouchMove), off(document, "touchmove", this._onTouchMove), off(document, "pointermove", this._onTouchMove), off(document, "dragover", nearestEmptyInsertDetectEvent), off(document, "mousemove", nearestEmptyInsertDetectEvent), off(document, "touchmove", nearestEmptyInsertDetectEvent);
  },
  _offUpEvents: function ze() {
    var Qr = this.el.ownerDocument;
    off(Qr, "mouseup", this._onDrop), off(Qr, "touchend", this._onDrop), off(Qr, "pointerup", this._onDrop), off(Qr, "pointercancel", this._onDrop), off(Qr, "touchcancel", this._onDrop), off(document, "selectstart", this);
  },
  _onDrop: function ze(Qr) {
    var Wr = this.el, Gr = this.options;
    if (newIndex = index(dragEl), newDraggableIndex = index(dragEl, Gr.draggable), pluginEvent("drop", this, {
      evt: Qr
    }), parentEl = dragEl && dragEl.parentNode, newIndex = index(dragEl), newDraggableIndex = index(dragEl, Gr.draggable), Sortable.eventCanceled) {
      this._nulling();
      return;
    }
    awaitingDragStarted = !1, isCircumstantialInvert = !1, pastFirstInvertThresh = !1, clearInterval(this._loopId), clearTimeout(this._dragStartTimer), _cancelNextTick(this.cloneId), _cancelNextTick(this._dragStartId), this.nativeDraggable && (off(document, "drop", this), off(Wr, "dragstart", this._onDragStart)), this._offMoveEvents(), this._offUpEvents(), Safari && css(document.body, "user-select", ""), css(dragEl, "transform", ""), Qr && (moved && (Qr.cancelable && Qr.preventDefault(), !Gr.dropBubble && Qr.stopPropagation()), ghostEl && ghostEl.parentNode && ghostEl.parentNode.removeChild(ghostEl), (rootEl === parentEl || putSortable && putSortable.lastPutMode !== "clone") && cloneEl && cloneEl.parentNode && cloneEl.parentNode.removeChild(cloneEl), dragEl && (this.nativeDraggable && off(dragEl, "dragend", this), _disableDraggable(dragEl), dragEl.style["will-change"] = "", moved && !awaitingDragStarted && toggleClass(dragEl, putSortable ? putSortable.options.ghostClass : this.options.ghostClass, !1), toggleClass(dragEl, this.options.chosenClass, !1), _dispatchEvent({
      sortable: this,
      name: "unchoose",
      toEl: parentEl,
      newIndex: null,
      newDraggableIndex: null,
      originalEvent: Qr
    }), rootEl !== parentEl ? (newIndex >= 0 && (_dispatchEvent({
      rootEl: parentEl,
      name: "add",
      toEl: parentEl,
      fromEl: rootEl,
      originalEvent: Qr
    }), _dispatchEvent({
      sortable: this,
      name: "remove",
      toEl: parentEl,
      originalEvent: Qr
    }), _dispatchEvent({
      rootEl: parentEl,
      name: "sort",
      toEl: parentEl,
      fromEl: rootEl,
      originalEvent: Qr
    }), _dispatchEvent({
      sortable: this,
      name: "sort",
      toEl: parentEl,
      originalEvent: Qr
    })), putSortable && putSortable.save()) : newIndex !== oldIndex && newIndex >= 0 && (_dispatchEvent({
      sortable: this,
      name: "update",
      toEl: parentEl,
      originalEvent: Qr
    }), _dispatchEvent({
      sortable: this,
      name: "sort",
      toEl: parentEl,
      originalEvent: Qr
    })), Sortable.active && ((newIndex == null || newIndex === -1) && (newIndex = oldIndex, newDraggableIndex = oldDraggableIndex), _dispatchEvent({
      sortable: this,
      name: "end",
      toEl: parentEl,
      originalEvent: Qr
    }), this.save()))), this._nulling();
  },
  _nulling: function ze() {
    pluginEvent("nulling", this), rootEl = dragEl = parentEl = ghostEl = nextEl = cloneEl = lastDownEl = cloneHidden = tapEvt = touchEvt = moved = newIndex = newDraggableIndex = oldIndex = oldDraggableIndex = lastTarget = lastDirection = putSortable = activeGroup = Sortable.dragged = Sortable.ghost = Sortable.clone = Sortable.active = null, savedInputChecked.forEach(function(Qr) {
      Qr.checked = !0;
    }), savedInputChecked.length = lastDx = lastDy = 0;
  },
  handleEvent: function ze(Qr) {
    switch (Qr.type) {
      case "drop":
      case "dragend":
        this._onDrop(Qr);
        break;
      case "dragenter":
      case "dragover":
        dragEl && (this._onDragOver(Qr), _globalDragOver(Qr));
        break;
      case "selectstart":
        Qr.preventDefault();
        break;
    }
  },
  /**
   * Serializes the item into an array of string.
   * @returns {String[]}
   */
  toArray: function ze() {
    for (var Qr = [], Wr, Gr = this.el.children, Yr = 0, Jr = Gr.length, Zr = this.options; Yr < Jr; Yr++)
      Wr = Gr[Yr], closest(Wr, Zr.draggable, this.el, !1) && Qr.push(Wr.getAttribute(Zr.dataIdAttr) || _generateId(Wr));
    return Qr;
  },
  /**
   * Sorts the elements according to the array.
   * @param  {String[]}  order  order of the items
   */
  sort: function ze(Qr, Wr) {
    var Gr = {}, Yr = this.el;
    this.toArray().forEach(function(Jr, Zr) {
      var Kr = Yr.children[Zr];
      closest(Kr, this.options.draggable, Yr, !1) && (Gr[Jr] = Kr);
    }, this), Wr && this.captureAnimationState(), Qr.forEach(function(Jr) {
      Gr[Jr] && (Yr.removeChild(Gr[Jr]), Yr.appendChild(Gr[Jr]));
    }), Wr && this.animateAll();
  },
  /**
   * Save the current sorting
   */
  save: function ze() {
    var Qr = this.options.store;
    Qr && Qr.set && Qr.set(this);
  },
  /**
   * For each element in the set, get the first element that matches the selector by testing the element itself and traversing up through its ancestors in the DOM tree.
   * @param   {HTMLElement}  el
   * @param   {String}       [selector]  default: `options.draggable`
   * @returns {HTMLElement|null}
   */
  closest: function ze(Qr, Wr) {
    return closest(Qr, Wr || this.options.draggable, this.el, !1);
  },
  /**
   * Set/get option
   * @param   {string} name
   * @param   {*}      [value]
   * @returns {*}
   */
  option: function ze(Qr, Wr) {
    var Gr = this.options;
    if (Wr === void 0)
      return Gr[Qr];
    var Yr = PluginManager.modifyOption(this, Qr, Wr);
    typeof Yr < "u" ? Gr[Qr] = Yr : Gr[Qr] = Wr, Qr === "group" && _prepareGroup(Gr);
  },
  /**
   * Destroy
   */
  destroy: function ze() {
    pluginEvent("destroy", this);
    var Qr = this.el;
    Qr[expando] = null, off(Qr, "mousedown", this._onTapStart), off(Qr, "touchstart", this._onTapStart), off(Qr, "pointerdown", this._onTapStart), this.nativeDraggable && (off(Qr, "dragover", this), off(Qr, "dragenter", this)), Array.prototype.forEach.call(Qr.querySelectorAll("[draggable]"), function(Wr) {
      Wr.removeAttribute("draggable");
    }), this._onDrop(), this._disableDelayedDragEvents(), sortables.splice(sortables.indexOf(this.el), 1), this.el = Qr = null;
  },
  _hideClone: function ze() {
    if (!cloneHidden) {
      if (pluginEvent("hideClone", this), Sortable.eventCanceled) return;
      css(cloneEl, "display", "none"), this.options.removeCloneOnHide && cloneEl.parentNode && cloneEl.parentNode.removeChild(cloneEl), cloneHidden = !0;
    }
  },
  _showClone: function ze(Qr) {
    if (Qr.lastPutMode !== "clone") {
      this._hideClone();
      return;
    }
    if (cloneHidden) {
      if (pluginEvent("showClone", this), Sortable.eventCanceled) return;
      dragEl.parentNode == rootEl && !this.options.group.revertClone ? rootEl.insertBefore(cloneEl, dragEl) : nextEl ? rootEl.insertBefore(cloneEl, nextEl) : rootEl.appendChild(cloneEl), this.options.group.revertClone && this.animate(dragEl, cloneEl), css(cloneEl, "display", ""), cloneHidden = !1;
    }
  }
};
function _globalDragOver(ze) {
  ze.dataTransfer && (ze.dataTransfer.dropEffect = "move"), ze.cancelable && ze.preventDefault();
}
function _onMove(ze, Qr, Wr, Gr, Yr, Jr, Zr, Kr) {
  var en, tn = ze[expando], nn = tn.options.onMove, rn;
  return window.CustomEvent && !IE11OrLess && !Edge ? en = new CustomEvent("move", {
    bubbles: !0,
    cancelable: !0
  }) : (en = document.createEvent("Event"), en.initEvent("move", !0, !0)), en.to = Qr, en.from = ze, en.dragged = Wr, en.draggedRect = Gr, en.related = Yr || Qr, en.relatedRect = Jr || getRect(Qr), en.willInsertAfter = Kr, en.originalEvent = Zr, ze.dispatchEvent(en), nn && (rn = nn.call(tn, en, Zr)), rn;
}
function _disableDraggable(ze) {
  ze.draggable = !1;
}
function _unsilent() {
  _silent = !1;
}
function _ghostIsFirst(ze, Qr, Wr) {
  var Gr = getRect(getChild(Wr.el, 0, Wr.options, !0)), Yr = getChildContainingRectFromElement(Wr.el, Wr.options, ghostEl), Jr = 10;
  return Qr ? ze.clientX < Yr.left - Jr || ze.clientY < Gr.top && ze.clientX < Gr.right : ze.clientY < Yr.top - Jr || ze.clientY < Gr.bottom && ze.clientX < Gr.left;
}
function _ghostIsLast(ze, Qr, Wr) {
  var Gr = getRect(lastChild(Wr.el, Wr.options.draggable)), Yr = getChildContainingRectFromElement(Wr.el, Wr.options, ghostEl), Jr = 10;
  return Qr ? ze.clientX > Yr.right + Jr || ze.clientY > Gr.bottom && ze.clientX > Gr.left : ze.clientY > Yr.bottom + Jr || ze.clientX > Gr.right && ze.clientY > Gr.top;
}
function _getSwapDirection(ze, Qr, Wr, Gr, Yr, Jr, Zr, Kr) {
  var en = Gr ? ze.clientY : ze.clientX, tn = Gr ? Wr.height : Wr.width, nn = Gr ? Wr.top : Wr.left, rn = Gr ? Wr.bottom : Wr.right, an = !1;
  if (!Zr) {
    if (Kr && targetMoveDistance < tn * Yr) {
      if (!pastFirstInvertThresh && (lastDirection === 1 ? en > nn + tn * Jr / 2 : en < rn - tn * Jr / 2) && (pastFirstInvertThresh = !0), pastFirstInvertThresh)
        an = !0;
      else if (lastDirection === 1 ? en < nn + targetMoveDistance : en > rn - targetMoveDistance)
        return -lastDirection;
    } else if (en > nn + tn * (1 - Yr) / 2 && en < rn - tn * (1 - Yr) / 2)
      return _getInsertDirection(Qr);
  }
  return an = an || Zr, an && (en < nn + tn * Jr / 2 || en > rn - tn * Jr / 2) ? en > nn + tn / 2 ? 1 : -1 : 0;
}
function _getInsertDirection(ze) {
  return index(dragEl) < index(ze) ? 1 : -1;
}
function _generateId(ze) {
  for (var Qr = ze.tagName + ze.className + ze.src + ze.href + ze.textContent, Wr = Qr.length, Gr = 0; Wr--; )
    Gr += Qr.charCodeAt(Wr);
  return Gr.toString(36);
}
function _saveInputCheckedState(ze) {
  savedInputChecked.length = 0;
  for (var Qr = ze.getElementsByTagName("input"), Wr = Qr.length; Wr--; ) {
    var Gr = Qr[Wr];
    Gr.checked && savedInputChecked.push(Gr);
  }
}
function _nextTick(ze) {
  return setTimeout(ze, 0);
}
function _cancelNextTick(ze) {
  return clearTimeout(ze);
}
documentExists && on(document, "touchmove", function(ze) {
  (Sortable.active || awaitingDragStarted) && ze.cancelable && ze.preventDefault();
});
Sortable.utils = {
  on,
  off,
  css,
  find,
  is: function ze(Qr, Wr) {
    return !!closest(Qr, Wr, Qr, !1);
  },
  extend,
  throttle,
  closest,
  toggleClass,
  clone,
  index,
  nextTick: _nextTick,
  cancelNextTick: _cancelNextTick,
  detectDirection: _detectDirection,
  getChild,
  expando
};
Sortable.get = function(ze) {
  return ze[expando];
};
Sortable.mount = function() {
  for (var ze = arguments.length, Qr = new Array(ze), Wr = 0; Wr < ze; Wr++)
    Qr[Wr] = arguments[Wr];
  Qr[0].constructor === Array && (Qr = Qr[0]), Qr.forEach(function(Gr) {
    if (!Gr.prototype || !Gr.prototype.constructor)
      throw "Sortable: Mounted plugin must be a constructor function, not ".concat({}.toString.call(Gr));
    Gr.utils && (Sortable.utils = _objectSpread2(_objectSpread2({}, Sortable.utils), Gr.utils)), PluginManager.mount(Gr);
  });
};
Sortable.create = function(ze, Qr) {
  return new Sortable(ze, Qr);
};
Sortable.version = version;
var autoScrolls = [], scrollEl, scrollRootEl, scrolling = !1, lastAutoScrollX, lastAutoScrollY, touchEvt$1, pointerElemChangedInterval;
function AutoScrollPlugin() {
  function ze() {
    this.defaults = {
      scroll: !0,
      forceAutoScrollFallback: !1,
      scrollSensitivity: 30,
      scrollSpeed: 10,
      bubbleScroll: !0
    };
    for (var Qr in this)
      Qr.charAt(0) === "_" && typeof this[Qr] == "function" && (this[Qr] = this[Qr].bind(this));
  }
  return ze.prototype = {
    dragStarted: function(Wr) {
      var Gr = Wr.originalEvent;
      this.sortable.nativeDraggable ? on(document, "dragover", this._handleAutoScroll) : this.options.supportPointer ? on(document, "pointermove", this._handleFallbackAutoScroll) : Gr.touches ? on(document, "touchmove", this._handleFallbackAutoScroll) : on(document, "mousemove", this._handleFallbackAutoScroll);
    },
    dragOverCompleted: function(Wr) {
      var Gr = Wr.originalEvent;
      !this.options.dragOverBubble && !Gr.rootEl && this._handleAutoScroll(Gr);
    },
    drop: function() {
      this.sortable.nativeDraggable ? off(document, "dragover", this._handleAutoScroll) : (off(document, "pointermove", this._handleFallbackAutoScroll), off(document, "touchmove", this._handleFallbackAutoScroll), off(document, "mousemove", this._handleFallbackAutoScroll)), clearPointerElemChangedInterval(), clearAutoScrolls(), cancelThrottle();
    },
    nulling: function() {
      touchEvt$1 = scrollRootEl = scrollEl = scrolling = pointerElemChangedInterval = lastAutoScrollX = lastAutoScrollY = null, autoScrolls.length = 0;
    },
    _handleFallbackAutoScroll: function(Wr) {
      this._handleAutoScroll(Wr, !0);
    },
    _handleAutoScroll: function(Wr, Gr) {
      var Yr = this, Jr = (Wr.touches ? Wr.touches[0] : Wr).clientX, Zr = (Wr.touches ? Wr.touches[0] : Wr).clientY, Kr = document.elementFromPoint(Jr, Zr);
      if (touchEvt$1 = Wr, Gr || this.options.forceAutoScrollFallback || Edge || IE11OrLess || Safari) {
        autoScroll(Wr, this.options, Kr, Gr);
        var en = getParentAutoScrollElement(Kr, !0);
        scrolling && (!pointerElemChangedInterval || Jr !== lastAutoScrollX || Zr !== lastAutoScrollY) && (pointerElemChangedInterval && clearPointerElemChangedInterval(), pointerElemChangedInterval = setInterval(function() {
          var tn = getParentAutoScrollElement(document.elementFromPoint(Jr, Zr), !0);
          tn !== en && (en = tn, clearAutoScrolls()), autoScroll(Wr, Yr.options, tn, Gr);
        }, 10), lastAutoScrollX = Jr, lastAutoScrollY = Zr);
      } else {
        if (!this.options.bubbleScroll || getParentAutoScrollElement(Kr, !0) === getWindowScrollingElement()) {
          clearAutoScrolls();
          return;
        }
        autoScroll(Wr, this.options, getParentAutoScrollElement(Kr, !1), !1);
      }
    }
  }, _extends(ze, {
    pluginName: "scroll",
    initializeByDefault: !0
  });
}
function clearAutoScrolls() {
  autoScrolls.forEach(function(ze) {
    clearInterval(ze.pid);
  }), autoScrolls = [];
}
function clearPointerElemChangedInterval() {
  clearInterval(pointerElemChangedInterval);
}
var autoScroll = throttle(function(ze, Qr, Wr, Gr) {
  if (Qr.scroll) {
    var Yr = (ze.touches ? ze.touches[0] : ze).clientX, Jr = (ze.touches ? ze.touches[0] : ze).clientY, Zr = Qr.scrollSensitivity, Kr = Qr.scrollSpeed, en = getWindowScrollingElement(), tn = !1, nn;
    scrollRootEl !== Wr && (scrollRootEl = Wr, clearAutoScrolls(), scrollEl = Qr.scroll, nn = Qr.scrollFn, scrollEl === !0 && (scrollEl = getParentAutoScrollElement(Wr, !0)));
    var rn = 0, an = scrollEl;
    do {
      var sn = an, ln = getRect(sn), un = ln.top, bn = ln.bottom, Cn = ln.left, hn = ln.right, gn = ln.width, dn = ln.height, yn = void 0, In = void 0, pn = sn.scrollWidth, vn = sn.scrollHeight, mn = css(sn), En = sn.scrollLeft, cn = sn.scrollTop;
      sn === en ? (yn = gn < pn && (mn.overflowX === "auto" || mn.overflowX === "scroll" || mn.overflowX === "visible"), In = dn < vn && (mn.overflowY === "auto" || mn.overflowY === "scroll" || mn.overflowY === "visible")) : (yn = gn < pn && (mn.overflowX === "auto" || mn.overflowX === "scroll"), In = dn < vn && (mn.overflowY === "auto" || mn.overflowY === "scroll"));
      var Sn = yn && (Math.abs(hn - Yr) <= Zr && En + gn < pn) - (Math.abs(Cn - Yr) <= Zr && !!En), xn = In && (Math.abs(bn - Jr) <= Zr && cn + dn < vn) - (Math.abs(un - Jr) <= Zr && !!cn);
      if (!autoScrolls[rn])
        for (var Dn = 0; Dn <= rn; Dn++)
          autoScrolls[Dn] || (autoScrolls[Dn] = {});
      (autoScrolls[rn].vx != Sn || autoScrolls[rn].vy != xn || autoScrolls[rn].el !== sn) && (autoScrolls[rn].el = sn, autoScrolls[rn].vx = Sn, autoScrolls[rn].vy = xn, clearInterval(autoScrolls[rn].pid), (Sn != 0 || xn != 0) && (tn = !0, autoScrolls[rn].pid = setInterval((function() {
        Gr && this.layer === 0 && Sortable.active._onTouchMove(touchEvt$1);
        var fn = autoScrolls[this.layer].vy ? autoScrolls[this.layer].vy * Kr : 0, Tn = autoScrolls[this.layer].vx ? autoScrolls[this.layer].vx * Kr : 0;
        typeof nn == "function" && nn.call(Sortable.dragged.parentNode[expando], Tn, fn, ze, touchEvt$1, autoScrolls[this.layer].el) !== "continue" || scrollBy(autoScrolls[this.layer].el, Tn, fn);
      }).bind({
        layer: rn
      }), 24))), rn++;
    } while (Qr.bubbleScroll && an !== en && (an = getParentAutoScrollElement(an, !1)));
    scrolling = tn;
  }
}, 30), drop = function ze(Qr) {
  var Wr = Qr.originalEvent, Gr = Qr.putSortable, Yr = Qr.dragEl, Jr = Qr.activeSortable, Zr = Qr.dispatchSortableEvent, Kr = Qr.hideGhostForTarget, en = Qr.unhideGhostForTarget;
  if (Wr) {
    var tn = Gr || Jr;
    Kr();
    var nn = Wr.changedTouches && Wr.changedTouches.length ? Wr.changedTouches[0] : Wr, rn = document.elementFromPoint(nn.clientX, nn.clientY);
    en(), tn && !tn.el.contains(rn) && (Zr("spill"), this.onSpill({
      dragEl: Yr,
      putSortable: Gr
    }));
  }
};
function Revert() {
}
Revert.prototype = {
  startIndex: null,
  dragStart: function ze(Qr) {
    var Wr = Qr.oldDraggableIndex;
    this.startIndex = Wr;
  },
  onSpill: function ze(Qr) {
    var Wr = Qr.dragEl, Gr = Qr.putSortable;
    this.sortable.captureAnimationState(), Gr && Gr.captureAnimationState();
    var Yr = getChild(this.sortable.el, this.startIndex, this.options);
    Yr ? this.sortable.el.insertBefore(Wr, Yr) : this.sortable.el.appendChild(Wr), this.sortable.animateAll(), Gr && Gr.animateAll();
  },
  drop
};
_extends(Revert, {
  pluginName: "revertOnSpill"
});
function Remove() {
}
Remove.prototype = {
  onSpill: function ze(Qr) {
    var Wr = Qr.dragEl, Gr = Qr.putSortable, Yr = Gr || this.sortable;
    Yr.captureAnimationState(), Wr.parentNode && Wr.parentNode.removeChild(Wr), Yr.animateAll();
  },
  drop
};
_extends(Remove, {
  pluginName: "removeOnSpill"
});
Sortable.mount(new AutoScrollPlugin());
Sortable.mount(Remove, Revert);
window.htmx = htmx;
window.htmx.onLoad((ze) => {
  var Qr = document.getElementById("items");
  Sortable.create(Qr, { handle: ".handle" });
});
