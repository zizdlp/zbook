"use client";
import { useEffect, useRef } from "react";
import * as d3 from "d3";
import { GeoProjection, GeoPath } from "d3";
import { FeatureCollection, GeoJsonProperties, Geometry } from "geojson";

interface Marker {
  ip: string;
  long: number;
  lat: number;
  city?: string;
  count: number;
}

export default function EarthChart({
  landData,
  lakeData,
  riverData,
  markers,
}: {
  landData: FeatureCollection<Geometry, GeoJsonProperties>;
  lakeData: FeatureCollection<Geometry, GeoJsonProperties>;
  riverData: FeatureCollection<Geometry, GeoJsonProperties>;
  markers: Marker[];
}) {
  markers = markers || [];
  const svgRef = useRef<SVGSVGElement>(null);

  useEffect(() => {
    const svg = d3.select(svgRef.current);
    const container = svg.node()?.parentElement;
    if (!container) return;
    const width = container.clientWidth;
    const height = container.clientHeight;
    svg.attr("width", width).attr("height", width);
    const tooltip = d3
      .select("body")
      .append("div")
      .style("position", "absolute")
      .style("background", "white")
      .style("border", "1px solid black")
      .style("padding", "5px")
      .style("border-radius", "5px")
      .style("font-size", "12px")
      .style("color", "black")
      .style("pointer-events", "none")
      .style("opacity", 0);

    const projection: GeoProjection = d3
      .geoOrthographic()
      .center([0, height > 600 ? -15 : -30])
      .scale(height > 600 ? 300 : 120)
      .clipAngle(90)
      .translate([width / 2, height / 2])
      .rotate([0, 0]);

    const path: GeoPath<any, any> = d3.geoPath().projection(projection);

    let currentRotation = [0, 0, 0];

    const drag = d3
      .drag()
      .subject(function () {
        const r = projection.rotate();
        return { x: r[0] / 0.5, y: -r[1] / 0.5 };
      })
      .on("drag", function (event) {
        const rotate = projection.rotate();
        projection.rotate([event.x * 0.5, -event.y * 0.5, rotate[2]]);
        svg.selectAll("path").attr("d", path);

        svg
          .selectAll<SVGCircleElement, Marker>("circle")
          .attr("cx", (d) => {
            const coords = projection([d.long, d.lat]);
            return coords ? coords[0] : 0;
          })
          .attr("cy", (d) => {
            const coords = projection([d.long, d.lat]);
            return coords ? coords[1] : 0;
          })
          .attr("display", (d) => {
            const distance = d3.geoDistance(
              [d.long, d.lat],
              [-projection.rotate()[0], -projection.rotate()[1]]
            );
            return distance > Math.PI / 2 ? "none" : "inline";
          });

        currentRotation = projection.rotate();
      });

    // 这里进行类型转换以解决类型不兼容问题
    (svg as unknown as d3.Selection<Element, unknown, null, undefined>).call(
      drag
    );

    // Draw land
    const landGroup = svg.append("g");
    landGroup
      .selectAll("path")
      .data(landData.features)
      .enter()
      .append("path")
      .attr("d", path)
      .style("fill", "#cce5ff")
      .style("stroke", "#333")
      .style("stroke-width", 0.5);

    // Draw lakes
    const lakeGroup = svg.append("g");
    lakeGroup
      .selectAll("path")
      .data(lakeData.features)
      .enter()
      .append("path")
      .attr("d", path)
      .style("fill", "#99ccff")
      .style("stroke", "#333")
      .style("stroke-width", 0.5);

    // Draw rivers
    const riverGroup = svg.append("g");
    riverGroup
      .selectAll("path")
      .data(riverData.features)
      .enter()
      .append("path")
      .attr("d", (d) => {
        const geometry = d.geometry;
        if (geometry.type === "LineString") {
          return path(geometry.coordinates);
        } else if (geometry.type === "MultiLineString") {
          return geometry.coordinates
            .map((coords: any) => path(coords))
            .join(" ");
        } else {
          return "";
        }
      })
      .style("fill", "none")
      .style("stroke", "#a8d6ff")
      .style("stroke-width", 1);

    // Draw markers
    svg
      .append("g")
      .selectAll("circle")
      .data(markers)
      .enter()
      .append("circle")
      .attr("cx", (d: Marker) => projection([d.long, d.lat])![0])
      .attr("cy", (d: Marker) => projection([d.long, d.lat])![1])
      .attr("r", 3)
      .style("fill", "brown")
      .on("mouseover", (event, d) => {
        tooltip
          .html(d.city ? d.city : d.ip)
          .style("left", `${event.pageX + 10}px`)
          .style("top", `${event.pageY + 10}px`)
          .style("opacity", 1);
      })
      .on("mousemove", (event) => {
        tooltip
          .style("left", `${event.pageX + 10}px`)
          .style("top", `${event.pageY + 10}px`);
      })
      .on("mouseout", () => {
        tooltip.style("opacity", 0);
      });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return <svg className="md:h-[655px] h-[330px]" ref={svgRef}></svg>;
}
